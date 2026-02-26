<template>
  <div id="app" ref="app">
    <div id="dashboard" ref="dashboard">
      <div id="errorModal" class="modal" v-show="showErrorModal">
        <p class="error-message">{{ errorMessage }} Please use this <a href="https://help.boldbi.com/site-administration/embed-settings/" target="_blank">link</a> to obtain the Json file from the Bold BI server.</p>
      </div>
    </div>
  </div>
</template>

<script>
import $ from 'jquery';
import { BoldBI } from '@boldbi/boldbi-embedded-sdk';
import axios from 'axios';

window.jQuery = $;

export default {
  name: 'App',
  data() {
    return {
      errorMessage: '',
    };
  },
  async mounted() {
    var scripts = [
      'https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js',
    ];
    scripts.forEach((script) => {
      let tag = document.createElement('script');
      tag.setAttribute('src', script);
      tag.setAttribute('type', 'text/javascript');
      tag.setAttribute('defer', 'defer');
      tag.async = true;
      document.head.appendChild(tag);
    });

    //Url of the tokenGeneration action in tokengeneration.go
    const tokenGenerationUrl = "http://localhost:8086/tokenGeneration";

    try {
      const response = await axios.get('http://localhost:8086/getdetails');
      if(response.data== null)
      {
        this.errorMessage = 'To compile and run the project, an embed config file needs to be required.';
        this.showErrorModal = true;
      }
      else{
        renderDashboard(response.data);
      } 
    } catch (error) {
        this.errorMessage = 'To compile and run the project, an embed config file needs to be required.';
        this.showErrorModal = true;
    }

    function getEmbedToken() {
          return fetch(tokenGenerationUrl, {  // Backend application URL
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({})
          })
            .then(response => {
              if (!response.ok) throw new Error("Token fetch failed");
              return response.text();
            });
        }
    
      function renderDashboard(data) {
        getEmbedToken()
          .then(accessToken => {
            const dashboard = BoldBI.create({
              serverUrl: data.ServerUrl + "/" + data.SiteIdentifier,
              dashboardId: data.DashboardId,
              embedContainerId: "dashboard",
              embedToken: accessToken
            });

            dashboard.loadDashboard();
          })
          .catch(err => {
            console.error("Error rendering dashboard:", err);
          });
      }
  }
};
</script>

<style>
.error-message {
  color: red;
  text-align: center;
  font-size: 20px;
  margin-top: 300px
}
</style>
