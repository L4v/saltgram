<template>
  <div style="background-color: #efeeee; height: 100vh">
    <portal-target name="drop-down-profile-menu" />
    <portal-target name="settings-menu"/>
    <TopBar style="position: sticky; z-index: 2"/>
    <div id="main-div"
         style="background-color: transparent">

      <!--        TODO: MILE-->
      <div id="notifications-header-div"
           style="background-color: white;">
        <h2 style="letter-spacing: 1px">Notifications</h2>
      </div>

      <div class="top-notification-bar mt-3" v-if="!privateProfile">
        <v-btn v-bind:class="NotificationCategory === 0 ? 'primary' : 'accent'"
               @click="NotificationCategory = 0"
               class="mx-2 my-1"
               small>
          Regular notifications
        </v-btn>
        <v-btn v-bind:class="NotificationCategory === 1 ? 'primary' : 'accent'"
               @click="NotificationCategory = 1"
               class="mx-2 my-1"
               small>
          Follow request
        </v-btn>
        <v-btn v-if="influencer"
        :class="NotificationCategory === 2 ? 'primary' : 'accent'"
        @click="NotificationCategory = 2"
        class="mx-2 my-1"
        small>
          Campaign requests
        </v-btn>
      </div>

      <div class="notifications-body-div" v-if="NotificationCategory === 0">

        <FollowNotification v-for="(item, index) in this.FollowNotification" 
          :key="index"
          :username-prop="item.referredUsername"
          :picture-prop="item.profilePictureURL"
          />

        <PostCommentNotification v-for="(item, index) in this.commentNotifications" 
          :key="index"
          :username-prop="item.referredUsername"
          :picture-prop="item.profilePictureURL"
          />

        <PostLikeNotification v-for="(item, index) in this.likeNotifications" 
          :key="index"
          :username-prop="item.referredUsername"
          :picture-prop="item.profilePictureURL"
          />

      </div>

      <div class="notifications-body-div" v-else-if="NotificationCategory === 1">

        <RequestProfile v-for="(item, index) in this.followingRequests" :key="index" :username-prop="item.username" :picture-prop="item.profilePicture" v-on:reload-requests="getFollowRequests()"/>

      </div>

      <div class="notifications-body-div" v-else-if="NotificationCategory === 2">
        <div id="requests">
          <div class="request" v-for="request in campaignRequests" :key="request.campaignId">
            <b>{{request.website}}</b>
            <v-spacer></v-spacer>
            <v-btn class="accept-btn" @click="acceptCampaign(request)">Accept</v-btn>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TopBar from "@/components/TopBar";
import CommentTagNotification from "@/components/notifications_components/CommentTagNotification";
import FollowNotification from "@/components/notifications_components/FollowNotification";
import PostCommentNotification from "@/components/notifications_components/PostCommentNotification";
import PostLikeNotification from "@/components/notifications_components/PostLikeNotification";
import PostTagNotification from "@/components/notifications_components/PostTagNotification";
import RequestProfile from "@/components/notifications_components/RequestProfile";

export default {
  name: "Notifications",
  components: {TopBar, CommentTagNotification, FollowNotification, PostCommentNotification, PostLikeNotification,
               PostTagNotification, RequestProfile},
  data: function () {
    return {
      campaignRequests: [],
      influencer: false,
      privateProfile: false,
      NotificationCategory: 0,
      followingRequests: [],
      notifications: [],
      likeNotifications: [],
      commentNotifications: [],
      FollowNotification: [],
    }
  },
  methods: {
    acceptCampaign: function(campaign) {
      this.refreshToken(this.getAHeader())
        .then(rr => {
          this.$store.state.jws = rr.data;
          let config = {
            headers: {
              'Authorization': 'Bearer ' + this.$store.state.jws,
              'Content-Type': 'text/html',
            },
          };
          this.axios.post("users/campaign", campaign.campaignId, config)
            .then(() => this.getCampaignRequests());
        })
    },

    getCampaignRequests: function() {
      this.refreshToken(this.getAHeader())
        .then(rr => {
          this.$store.state.jws = rr.data;
          this.axios.get("users/get/campaign", {headers: this.getAHeader()})
            .then(r => this.campaignRequests = r.data);
        })
    },
    isInfluencer: function() {
      if(!this.$store.state.jws) {
        return;
      }
      this.refreshToken(this.getAHeader())
        .then(rr => {
          this.$store.state.jws = rr.data;
          this.axios.get('users/get/isinfluencer', {headers: this.getAHeader()})
            .then(() => {
              this.influencer = true;
              this.getCampaignRequests();
            })
            .catch(() => this.influencer = false);
        })
    },
    getFollowRequests: function() {
      this.axios.get("users/follow/requests/", {headers: this.getAHeader()})
      .then(r => {
        this.followingRequests = r.data;
        console.log(r.data);
        this.getNotifications();
      }).catch(err => {
        console.log(err);
      })
    },
    getNotifications: function() {
      this.axios.get("notification/", {headers: this.getAHeader()})
      .then(r => {
          console.log(r)
          this.notifications = r.data
          this.notifications.forEach(element => {
            if(element.type == "LIKE") {
              this.likeNotifications.push(element);
            } else if(element.type == "COMMENT") {
              this.commentNotifications.push(element)
            } else if(element.type == "FOLLOW") {
              this.FollowNotification.push(element);
            }
          })
          console.log(this.likeNotifications);
      });
    }
  },
  mounted() {
    this.isInfluencer();
    this.getFollowRequests();
  }
}
</script>

<style scoped>

#main-div {
  display: inline-block;
  margin: 15px 20% 0 20%;
  width: 60vw;
  height: 85vh;
  flex-direction: row;
  justify-content: center;
  align-content: center;


}

#notifications-header-div {
  display: flex;
  height: auto;
  flex-direction: column;
  text-align: -webkit-center;
  border: black 1px solid;

  border-start-end-radius: 10px 10px;
  border-end-end-radius: 10px 10px;
  border-start-start-radius: 10px 10px;
  border-end-start-radius: 10px 10px;

}

.notifications-body-div {
  background-color: #FFFFFF;
  height: 65vh;
  overflow-x: hidden;
  overflow-y: scroll;
  margin-top: 15px;
  /*display: flex;*/
  /*height: auto;*/
  /*flex-direction: column;*/
  /*text-align: -webkit-center;*/

  border: black 2px solid;

  border-start-end-radius: 10px 10px;
  border-end-end-radius: 10px 10px;
  border-start-start-radius: 10px 10px;
  border-end-start-radius: 10px 10px;

}

.top-notification-bar {
  position: static;
  display: inline-flex;
  padding: 0 25%;
  width: 100%;
  height: 40px;
  background-color: #FFFFFF;
  text-align: -webkit-center;

  border: 1px solid black;
  border-start-end-radius: 10px 10px;
  border-end-end-radius: 10px 10px;
  border-start-start-radius: 10px 10px;
  border-end-start-radius: 10px 10px;

}

.accept-btn {
  color: #26a900;
  border-color: #26a900;
}

.request {
  height: 70px;
  width: 100%;

  display: flex;
  flex-direction: row;
  flex-flow: wrap;
  transition: 0.3s;
  cursor: pointer;
  align-items: center;
  padding-left: 10px;
  padding-right: 10px;
}

.request:hover {
  background-color: #e0e0e0;
  transition: 0.3s;
}

</style>