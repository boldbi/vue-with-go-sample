<template>
  <div id="app" ref="app">
    <div id="dashboard" ref="dashboard">
      <div id="errorModal" class="modal" v-show="showErrorModal">
        <p class="error-message">{{ errorMessage }} You can get the JSON file from the Bold BI server using this <a v-if="showErrorModal" href="https://help.boldbi.com/embedded-bi/site-administration/embed-settings/" target="_blank">link</a>.</p>
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

    let authorizationUrl = 'http://localhost:8086/authorizationserver';

    try {
      
      const response = await axios.get('http://localhost:8086/getdetails');
       if(response.data== null)
      {
        this.errorMessage = 'To compile and run the project, an embed config file needs to be required.';
        this.showErrorModal = true;
      }
      else{
        createBoldBIDashboard(response.data);
      }
      
    } catch (error) {
        this.errorMessage = 'To compile and run the project, an embed config file needs to be required.';
        this.showErrorModal = true;
    }

    function createBoldBIDashboard(data) {
      let dashboard = BoldBI.create({
        serverUrl: data.ServerUrl + '/' + data.SiteIdentifier,
        dashboardId: data.DashboardId,
        embedContainerId: 'dashboard',
        embedType: BoldBI.EmbedType.Component,
        environment: data.Environment,
        width: '100%',
        height: window.innerHeight + 'px',
        expirationTime: 100000,
        authorizationServer: {
          url: authorizationUrl,
        },
      });
      dashboard.loadDashboard();
    }
  },
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
