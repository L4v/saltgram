<template>
  <div id="user-main">
    <portal-target name="drop-down-profile-menu" />
    <portal-target name="settings-menu"/>
    <TopBar style="position: sticky; z-index: 2"/>
    <UserNotFound v-if="!active" />
    <div id="user-header" v-if="this.active">
      <div id="user-icon-logout">
        <v-layout align-center
                  justify-center>
          <ProfileImage ref="profileImage"
                        :following-prop="this.followingUser"
                        :is-my-profile-prop="this.isMyProfile"
                        :username="this.profile.username"
                        :verified="profile.verified"
                        v-bind:image-src="profile.profilePictureURL"
                        @toggle-following="toggleFollow"/>
        </v-layout>
        <v-layout column
                  style="width: 70%"
                  justify-space-between>
          <v-layout style="height: 40%;"
                    justify-center
                    column>

            <ProfileHeader :following-prop="this.profile.following"
                           :followers-prop="this.profile.followers"
                           :user-prop="this.loggedUser.username"
                           :posts-number="usersPosts.length"/>

          </v-layout>

          <NameAndDescription :name="this.profile.fullName"
                              :description="this.profile.description"
                              :web-site="this.profile.webSite"
                              :account-type="this.profile.accountType"
                              @influencer="influencerRequest"/>
          <v-dialog v-model="influencerDialog" width="600">
            <div id="campaigns">
              <div v-for="campaign in campaigns" :key="campaign.id" class="campaign" @click="requestInfluencer(campaign)">
                <h3>{{campaign.website}}</h3>
                <v-spacer></v-spacer>
                <img v-if="campaign.url" :src="campaign.url" />
              </div>
            </div>
          </v-dialog>

        </v-layout>

      </div>
    </div>

    <div v-if="!isContentVisible && active" class="private-account">
      <i class="fa fa-lock" style="transform: scale(2.5)"/>
      <h3>This user is private</h3>

    </div>

<!--        TODO: STORY HIGHLIGHTS-->
    <v-layout id="user-stories"
              v-if="isContentVisible && active"
              column>
      <v-layout class="inner-story-layout"
                style="margin: 10px">
        <StoryHighlight v-for="highlight in highlights" :key="highlight.name" :stories="highlight.stories" :name="highlight.name"/>
        <div style="cursor: pointer" v-if="isMyProfile" id="new-highlight" @click="openHighlightDialog">
          +
        </div>
        <v-dialog
        v-model="highlightDialog"
        width="500px">
          <div v-if="highlightSuccess" class="success-dialog">
            <p><i class="fa fa-check" aria-hidden="true"></i></p>
            <b>Highlight added!</b>
          </div>
          <v-card v-else>
            <v-card-title>Add highlight</v-card-title>
            <v-card-text>
              <v-form v-model="highlightForm">
                <v-text-field
                v-model="highlightName"
                label="Highlight name"
                :rules="[noempty]"
                required/>
                <div id="story-selection">
                  <div v-for="(story, i) in stories" :key="story.filename" class="story-thumbnail" @click="selectStory(i)">
                    <v-img :src="story.url" width="128px" height="128px" />
                    <v-simple-checkbox class="story-checkbox" v-model="stories[i].isSelected" absolute @click="selectStory(i)" />
                  </div>
                </div>
              </v-form>
            </v-card-text>
            <v-card-action>
              <div class="d-flex flex-row">
                <v-btn plain @click="highlightDialog = false">Cancel</v-btn>
                <v-spacer></v-spacer>
                <v-btn :disabled="!highlightForm" color="accent" @click="addHighlight">Add highlight</v-btn>
              </div>
            </v-card-action>
          </v-card>
        </v-dialog>
      </v-layout>
    </v-layout >

    <!--  TODO: LAYOUT FOR TOGGLING: POSTS, SAVED, TAGGED  -->
    <v-layout id="radio-button-layout"
              v-if="isContentVisible && active">
      <v-radio-group row  v-model="radioButton">
        <v-radio label="Posts"  value="posts"/>
        <v-radio label="Saved"  value="saved"/>
        <v-radio label="Tagged" value="tagged"/>
      </v-radio-group>
    </v-layout>

<!--        TODO: POSTS -->
    <transition name="fade">
      <v-layout class="user-media"
                v-if="radioButton === 'posts' && isContentVisible && active"
                column>
                <div v-for="(object, index) in usersPosts" :key="index">
                  <PostOnUserPage :post="object"
                                  :user="{username: profile.username, profilePictureURL: profile.profilePictureURL}"/>
                </div>
      </v-layout>
    </transition>


    <!--        TODO: SAVED -->
    <transition name="fade">
      <v-layout class="user-media"
                v-if="radioButton === 'saved' && isContentVisible && active"
                column>
                <div v-for="(object, index) in savedPosts" :key="index">
                  <PostOnUserPage :post="object"
                                  :user="{username: object.user.username, profilePictureURL: object.user.profilePictureURL}"/>
                </div>
      </v-layout>
    </transition>

    <!--        TODO: TAGGED -->
    <transition name="fade">
      <v-layout class="user-media"
                v-if="radioButton === 'tagged' && isContentVisible && active"
                column>
                <div v-for="(object, index) in taggedPosts" :key="index">
                  <PostOnUserPage :post="object"
                                  :user="{username:  object.user.username, profilePictureURL:  object.user.profilePictureURL}"/>
                </div>
      </v-layout>
    </transition>
  </div>
</template>

<script>
import TopBar from "@/components/TopBar";
import ProfileImage from "@/components/user_page_components/ProfileImage";
import ProfileHeader from "@/components/user_page_components/ProfileHeader";
import NameAndDescription from "@/components/user_page_components/NameAndDescription";
import StoryHighlight from "@/components/user_page_components/StoryHighlight";
import PostOnUserPage from "@/components/user_page_components/PostOnUserPage";
import UserNotFound from "@/components/user_page_components/UserNotFound";

export default {
    components: {
      TopBar, ProfileImage, ProfileHeader, NameAndDescription, StoryHighlight, PostOnUserPage, UserNotFound
    },
    data: function() {
      return {
        campaigns: [
          {
            id: 1,
            website: 'web.com',
            url: 'https://icatcare.org/app/uploads/2018/07/Thinking-of-getting-a-cat.png',
          },
        ],
        influencerDialog: false,
        highlightSuccess: false,
        noempty: v => !!v || 'Required',
        highlightName: '',
        highlightForm: false,
        stories: [],
        highlightDialog: false,
        highlights: [],
        //
        user: {},
        profile : {
          privateUser: true,
          description: '',
          fullName: '',
          followers: 0,
          following: 0,
          followersList:[],
          followingList: [],
          username: '',
          webSite: '',
          profilePictureURL: '',
          verified: false,
          accountType: ''
        },
        isMyProfile: false,
        radioButton: 'posts',
        followingUser: false,
        pendingRequest: false,
        usersPosts: [],
        userStories: [],
        loggedUser: { username: '' },
        taggedPosts: [],
        savedPosts: [],
        active: true,
      }
    },
    computed: {
      isContentVisible: function() {
          return !(!this.isMyProfile && this.profile.privateUser && !this.followingUser);

      }
    },
    methods: {

      getCampaigns: function() {
        if(!this.$store.state.jws) {
          return;
        }
        this.refreshToken(this.getAHeader())
          .then(rr => {
            this.$store.state.jws = rr.data;
            this.axios.get("content/campaigns", {headers: this.getAHeader()})
              .then(r => this.campaigns = r.data)
          })
      },

      influencerRequest: function() {
        if(!this.followingUser && this.profile.privateUser) {
          return;
        }

        this.influencerDialog = true;
      },

      requestInfluencer: function(campaign) {
        this.refreshToken(this.getAHeader())
          .then(rr => {
            this.$store.state.jws = rr.data;
            let request = {
              influencerId: this.user.id,
              campaignId: campaign.id,
              website: campaign.website,
            };
            this.axios.post('users/influencer', request, {headers: this.getAHeader()})
              .then(() => this.influencerDialog = false)
          })
      },

        getHighlights: function() {
          this.refreshToken(this.getAHeader())
            .then(rr => {
              this.$store.state.jws = rr.data;
              this.axios.get('content/highlight/' + this.user.id)
                .then(r => this.highlights = r.data)
                .catch(r => console.log(r));
            }).catch(r => console.log(r));
        },

        addHighlight: function() {
          this.highlightSuccess = false;
          let data = {
            name: this.highlightName,
            stories: [],
          };
          this.stories.forEach(s => {
            if(s.isSelected) {
              data.stories.push(s);
            }
          });

          this.refreshToken(this.getAHeader())
            .then(rr => {
              this.$store.state.jws = rr.data;
              this.axios.post('content/highlight', data, {headers: this.getAHeader()})
                .then(() => this.highlightSuccess = true)
                .catch(r => console.log(r));
            }).catch(() => this.$router.push('/'));

        },

        selectStory: function(index) {
          this.stories[index].isSelected = !this.stories[index].isSelected;
          this.stories = [...this.stories];
        },

        openHighlightDialog: function() {
          this.highlightDialog = true;
          this.refreshToken(this.getAHeader())
            .then(rr => {
              this.$store.state.jws = rr.data;
              this.axios.get('content/story/' + this.user.id)
                .then(r => {
                  this.stories = [];
                  r.data.forEach(s => {
                    s.stories.forEach(ss => {
                      let newSS = ss;
                      newSS.closeFriends = s.closeFriends;
                      this.stories.push(newSS);
                    });
                  });
                  this.stories.forEach(s => {
                    s.isSelected = false;
                  })
                })
            }).catch(() => this.$router.push('/'));
        },

        getUserInfo: function() {
            // this.refreshToken(this.getAHeader())
            //     .then(rr => {
            //         this.$store.state.jws = rr.data;
            //         this.axios.get('/users', {headers: this.getAHeader()})
            //         .then(r => {
            //           this.isMyProfile = r.data.isThisMe;
            //         })
                    this.axios.get("users/" + this.$route.params.username)
                        .then(r =>{ 
                          this.user = r.data
                          this.getHighlights();
                          this.getUser();
                          });
                      
            //     }).catch(() => {
            //   console.log('No User is logged in!');
            // });
        },

        getUser: function() {
            // this.isMyProfile = this.user.username === this.$route.params.username;

            this.followingUser = false;
            this.privateUser = true;


            this.axios.get("users/profile/" + this.$route.params.username, {headers: this.getAHeader()})
            .then(r => {
              this.profile.privateUser = !r.data.isPublic;
              this.profile.followingUser = r.data.isFollowing;
              this.profile.username = r.data.username;
              this.profile.followers = r.data.followers;
              this.profile.following = r.data.following;
              this.profile.fullName = r.data.fullName;
              this.profile.description = r.data.description;
              this.profile.webSite = r.data.webSite;
              this.profile.profilePictureURL = r.data.profilePictureURL;
              this.profile.verified = r.data.verified;
              this.profile.accountType = r.data.accountType;
              this.isMyProfile = r.data.isThisMe;
              this.$refs.profileImage.$data.isMyProfile = this.isMyProfile
              // console.log(r.data.userId)
              // console.log('this.profile.followers: ', this.profile.fol)
              this.getUserPosts(r.data.userId);
              this.getUserStories(r.data.userId);
              this.getTaggedPosts();
              this.getSavedPosts();
            }).catch(err => {
              console.log(err)
            })
            if(!this.isMyProfile) {
              this.axios.get("users/check/follow/" + this.$route.params.username, {headers: this.getAHeader()})
              .then(r => {
                this.followingUser = r.data
                this.$refs.profileImage.$data.following = r.data
              })
        
              if(!this.followingUser) {
                this.axios.get("users/check/followrequest/" + this.$route.params.username, {headers: this.getAHeader()})
                .then(r => {
                  //this.pendingRequest = r.data
                  this.$refs.profileImage.$data.waitingForResponse = r.data
                })
              }
              this.checkIfMuted();
              this.checkIfBlocked();
            }
            
        },
        getUserPosts(id) {
           this.axios.get("content/post/" + id, {headers: this.getAHeader()})
           .then(r => {
               this.usersPosts = r.data;
               this.usersPosts.sort(function (a,b) {
                 let index1 = a.post.sharedMedia.media[0].addedOn.indexOf('CEST') + 4
                 let index2 = b.post.sharedMedia.media[0].addedOn.indexOf('CEST') + 4
                 let d1 = new Date(a.post.sharedMedia.media[0].addedOn.substring(0, index1).replace('CEST', '(CEST)'))
                 let d2 = new Date(b.post.sharedMedia.media[0].addedOn.substring(0, index2).replace('CEST', '(CEST)'))
                 if (d1 < d2) {
                   return 1;
                 }
                 if (d1 > d2) {
                   return -1;
                 }
                 // dates must be equal
                 return 0;
               });
              // console.log(this.usersPosts);
            }).catch(err => {
              console.log(err)
              //this.$router.push('/');
            })
        },
        getTaggedPosts() {
           this.axios.get("content/taggedposts" /*+ id*/, {headers: this.getAHeader()})
           .then(r => {
             console.log(r.data);
               this.taggedPosts = r.data;
               this.taggedPosts.sort(function (a,b) {
                 let index1 = a.post.sharedMedia.media[0].addedOn.indexOf('CEST') + 4
                 let index2 = b.post.sharedMedia.media[0].addedOn.indexOf('CEST') + 4
                 let d1 = new Date(a.post.sharedMedia.media[0].addedOn.substring(0, index1).replace('CEST', '(CEST)'))
                 let d2 = new Date(b.post.sharedMedia.media[0].addedOn.substring(0, index2).replace('CEST', '(CEST)'))
                 if (d1 < d2) {
                   return 1;
                 }
                 if (d1 > d2) {
                   return -1;
                 }
                 // dates must be equal
                 return 0;
               });
              // console.log(this.usersPosts);
            }).catch(err => {
              console.log(err)
              //this.$router.push('/');
            })
        },
        getSavedPosts() {
           this.axios.get("content/savedposts" /*+ id*/, {headers: this.getAHeader()})
           .then(r => {
                          console.log(r.data);
               this.savedPosts = r.data;
               this.savedPosts.sort(function (a,b) {
                 let index1 = a.post.sharedMedia.media[0].addedOn.indexOf('CEST') + 4
                 let index2 = b.post.sharedMedia.media[0].addedOn.indexOf('CEST') + 4
                 let d1 = new Date(a.post.sharedMedia.media[0].addedOn.substring(0, index1).replace('CEST', '(CEST)'))
                 let d2 = new Date(b.post.sharedMedia.media[0].addedOn.substring(0, index2).replace('CEST', '(CEST)'))
                 if (d1 < d2) {
                   return 1;
                 }
                 if (d1 > d2) {
                   return -1;
                 }
                 // dates must be equal
                 return 0;
               });
              // console.log(this.usersPosts);
            }).catch(err => {
              console.log(err)
              //this.$router.push('/');
            })
        },
        getUserStories(id) {
           this.axios.get("content/story/" + id, {headers: this.getAHeader()})
           .then(r => {
              //console.log(JSON.parse(r.data.toString()));
              this.userStories = r.data;
              console.log("stories:", r.data);
             const oneDay = 60 * 60 * 24 * 1000;
              let newStories = []
              if (this.userStories)  {
                this.userStories.forEach(s => {
                  s.stories.forEach(ss => {
                    let newSS = ss;
                    newSS.closeFriends = s.closeFriends;
                    newSS.isCampaign = s.isCampaign;
                    newSS.campaignWebsite = s.campaignWebsite;
                    let storyDate = new Date(Date.parse(newSS.addedOn.substring(0, 10)));
                    if ((Date.now() - storyDate) < oneDay && !s.isCampaign) newStories.push(newSS);
                    if(s.isCampaign) newStories.push(newSS);
                  });
                });
               this.$refs.profileImage.$data.userStories = newStories;//this.userStories;
              }
              

            }).catch(err => {
              console.log(err)
            //  this.$router.push('/');
            })
        },

        toggleFollow(follow) {
          this.followingUser = follow
        },
        getFollowingNumb() {
          this.getUserInfo();
        },

        getLoggedUserInfo() {
          this.axios.get('users', {headers: this.getAHeader()})
              .then(r => {
                this.loggedUser = r.data
              })
              .catch(err => {
                console.log(err)
              })
        },
        checkIfMuted: function() {  
          if(!this.isMyProfile){
            this.axios.get("users/check/muted/" + this.$route.params.username, {headers: this.getAHeader()})
            .then(r => {
              this.$refs.profileImage.$data.muted = r.data;
            })
            .catch(r => {
              console.log(r);
            })
          }
        },
        checkIfBlocked: function() {  
          if(!this.isMyProfile){
            this.axios.get("users/check/blocked/" + this.$route.params.username, {headers: this.getAHeader()})
            .then(r => {
              this.$refs.profileImage.$data.blocked = r.data;
            })
            .catch(r => {
              console.log(r);
            })
          }
        },
    },
    mounted() {
      this.getCampaigns();
        this.axios.get("users/check/" + this.$route.params.username)
        .then(r => {
          this.active = r.data
        })
        .catch(r => {
          this.active = false
          console.log(r)
        })
        this.axios.get("users/check/block/" + this.$route.params.username, {headers: this.getAHeader()})
          .then(r => {
            if (r.data) {
              //this.$router.push('/main/')
              this.active = false
            }
          })
       this.getUserInfo(); // TODO UNCOMMENT THIS
       this.getLoggedUserInfo();
    },
}
</script>

<style scoped>

    #user-main {
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-content: center;
        /* text-align: center; */
        background: #efeeee;
        min-height: 100vh;
        height: auto;
    }

    #user-header {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        /*background: cadetblue;*/
        margin-left: 10%;
        margin-right: 10%;
        height: auto;
    }

    #user-icon-logout {
        display: inline-flex;
        flex-direction: row;
        justify-content: center;
        /*background-color: red;*/
        width: 100%;
      height: auto;
    }

    #user-stories {
        flex-direction: column;
        height: auto;
        margin: 0 10%;
    }

    .user-media {
      --w:400px;
      --n:3;
      --m:2;

      margin: 0 10%;
      display:grid;
      grid-template-columns:repeat(auto-fit,minmax(clamp(100%/(var(--n) + 1) + 0.1%,(var(--w) - 100vw)*1000,100%/(var(--m) + 1) + 0.1%),1fr)); /*this */
      gap:10px;
    }

    #radio-button-layout {
      height: 70px;
      text-align: -webkit-center;
      justify-content: center;
      margin: 0 10%;
      float: left;
     }

    .inner-story-layout {
      height: 150px;
      flex-direction: row;
      overflow-x: auto;
      overflow-y: hidden;
      white-space: nowrap;
      align-items: center;
    }

    .inner-post-layout > div {
      display: inline-block;
      color: white;
      text-align: center;
      padding: 14px;
      text-decoration: none;
    }

    .private-account {
      text-align: -webkit-center;
      padding-top: 50px;
      height: 100px;
    }

    #new-highlight {
      display: grid;
      place-items: center;
      width: 80px;
      height: 80px;
      background: #ddd;
      font-weight: 500;
      font-size: 3rem;
    }

    #story-selection {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
    }

    .story-thumbnail {
      cursor: pointer;
      position: relative;
      display: inline-block;
    }

    .story-checkbox {
      position: absolute;
      top: 0;
      right: 15px;
    }

    .success-dialog {
      display: flex;
      flex-direction: column;
      justify-content: center;
      text-align: center;
      background: #fff;
      min-height: 40vh;
    }

    .success-dialog p {
      font-size: 4rem;
    }

    #campaigns {
      background: #fff;
      height: 400px;
      overflow: auto;
      padding: 10px;
    }

    .campaign {
      cursor: pointer;
      display: flex;
      flex-direction: row;
      padding: 10px;
      border-bottom: solid 1px #eee;
    }

    .campaign:hover {
      background: #eee;
    }

    .campaign img {
      height: 128px;
      width: 128px;
    }

</style>
