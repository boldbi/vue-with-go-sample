# BoldBI Embedding Vue with Go Sample

This Bold BI Vue with Go sample contains the Dashboard embedding sample.In this sample, the Vue application acts as the front-end, and the Go sample act as the back-end application.This sample demonstrates the dashboard rendering with the available dashboard in your Bold BI server.

This section guides you using the Bold BI dashboard in your Vue with Go sample application.

 * [Requirements to run the demo](#requirements-to-run-the-demo)
 * [Using the Vue with Go sample](#using-the-vue-with-go-sample)
 * [Online Demos](#online-demos)
 * [Documentation](#documentation)

 ## Requirements to run the demo

The samples require the following requirements to run

 * [Go installer](https://go.dev/dl/)
 * [Visual Studio Code](https://code.visualstudio.com/download)
 * [Node.js](https://nodejs.org/en/)

 ## Using the Vue with Go sample
 
 * Open the file `main.go` of the Go sample in Visual studio code. 

 * Please change the following properties in the `main.go` file as per your Bold BI Server.

<meta charset="utf-8"/>
<table>
  <tbody>
    <tr>
        <td align="left">EmbedSecret</td>
        <td align="left">Get your EmbedSecret key from the Embed tab by enabling the `Enable embed authentication` on the Administration page https://help.boldbi.com/embedded-bi/site-administration/embed-settings/.</td>
    </tr>
    <tr>
        <td align="left">UserEmail</td>
        <td align="left">UserEmail of the Admin in your Bold BI, which would be used to get the dashboard list.</td>
    </tr>
  </tbody>
</table>

* Now run the back-end Go sample by using the following command in the terminal.

```bash
 go run main.go
```

* Open the `Vue` sample in a new window of Visual studio code.

* Open the `App.Vue` in the following location, /src/App.Vue.

* Please change the following properties in the `App.Vue` file as per your Bold BI server and back-end application.

    <meta charset="utf-8"/>
    <table>
    <tbody>
        <tr>
            <td align="left">rootUrl</td>
            <td align="left">Dashboard Server URL (Eg: http://localhost:5000/bi, http://demo.boldbi.com/bi).</td>
        </tr>
        <tr>
            <td align="left">siteIdentifier</td>
            <td align="left">For the Bold BI Enterprise edition, it should be like `site/site1`. For Bold BI Cloud, it should be an empty string.</td>
        </tr>
        <tr>
            <td align="left">authorizationUrl</td>
            <td align="left">Url of the GetDetails action in the Go application(http://localhost:8086/getDetails).</td>
        </tr>
        <tr>
            <td align="left">environment</td>
            <td align="left">Your Bold BI application environment. (If Cloud, you should use `cloud,` if Enterprise, you should use `onpremise`).</td>
        </tr>
        <tr>
            <td align="left">dashboardId</td>
            <td align="left">Provide the dashboard id of the dashboard you want to embed in view or edit mode. Ignore this property to create a new dashboard.</td>
        </tr>
    </tbody>
    </table>



### Install npm

To install all dependent packages, use the below command

```bash
npm install
```

### Install Bold BI Embedded SDK package

To install the Bold BI Embedded SDK package, use the following command

```bash
npm install -save @boldbi/boldbi-embedded-sdk
```

### Run/Serve

To run the sample, use the below command

```bash
npm run serve
```
Runs the app in the development mode.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />

Please refer to the [help documentation](https://help.boldbi.com/embedded-bi/javascript-based/samples/v3.3.40-or-later/vuejs-with-go/#how-to-run-the-sample) to know how to run the sample.

## Online Demos

Look at the Bold BI Embedding sample to live demo [here](https://samples.boldbi.com/embed).


## Documentation

A complete Bold BI Embedding documentation can be found on [Bold BI Embedding Help](https://help.boldbi.com/embedded-bi/javascript-based/).


