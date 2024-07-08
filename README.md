# BoldBI Embedding Vue with Go Sample

This Bold BI Vue with Go sample contains the Dashboard embedding sample. In this sample, the Vue application acts as the front-end, and the Go sample act as the back-end application.This sample demonstrates the dashboard rendering with the available dashboard in your Bold BI server.

## Dashboard view

   ![Dashboard view](https://github.com/boldbi/vue-with-go-sample/assets/129486688/381aa89c-6870-4489-a744-c3617abc7646)

## Prerequisites

 * [Go installer](https://go.dev/dl/)
 * [Visual Studio Code](https://code.visualstudio.com/download)
 * [Node.js](https://nodejs.org/en/)
 > **NOTE:** Node.js v12.13 to v20.15 are supported.

 #### Supported browsers
  
  * Google Chrome, Microsoft Edge, Mozilla Firefox.

## Configuration

 * Please ensure that you have enabled embed authentication on the `embed settings` page. If it is not currently enabled, please refer to the following image or detailed [instructions](https://help.boldbi.com/site-administration/embed-settings/#get-embed-secret-code) to enable it.

    ![Embed Settings](https://github.com/boldbi/aspnet-core-sample/assets/91586758/b3a81978-9eb4-42b2-92bb-d1e2735ab007)

 * To download the `embedConfig.json` file, please follow this [link](https://help.boldbi.com/site-administration/embed-settings/#get-embed-configuration-file) for reference. Additionally, you can refer to the following image for visual guidance.

    ![Embed Settings Download](https://github.com/boldbi/aspnet-core-sample/assets/91586758/d27d4cfc-6a3e-4c34-975e-f5f22dea6172)
    ![EmbedConfig Properties](https://github.com/boldbi/aspnet-core-sample/assets/91586758/d6ce925a-0d4c-45d2-817e-24d6d59e0d63)

 * Copy the downloaded `embedConfig.json` file and paste it into the designated [location](https://github.com/boldbi/vue-with-go-sample/tree/master/Go) within the application. Please ensure that you have placed it in the application as shown in the following image.

   ![EmbedConfig image](https://github.com/boldbi/vue-with-go-sample/assets/129486688/bf994470-ed88-46e3-9b3e-2c941d42a2a6)
v
 ## Run a Sample Using Command Line Interface

  1. Open the **command line interface** and navigate to the specified file [location](https://github.com/boldbi/vue-with-go-sample/tree/master/Go) where the project is located.
   
  2. Run the back-end `Go` sample by using the following command `go run main.go`.
   
  3. Open the **command line interface** and navigate to the specified file [location](https://github.com/boldbi/vue-with-go-sample/tree/master/Vue) where the project is located.
   
  4. Install all dependent packages by executing the following command `npm install`.
   
  5. Finally, run the application using the following command `npm run serve`.
   
  6. After the application has started, it will display a URL in the `command line interface`, typically something like (e.g., https://localhost:8080). Copy this URL and paste it into your default web browser.

 ## Developer IDE

  * [Visual studio code](https://code.visualstudio.com/download)

 ### Run a Sample Using Visual Studio Code

 * Open the `Go` sample in **Visual Studio Code.**

 * Run the back-end `Go` sample by using the following command in the terminal `go run main.go`.

 * Open the `Vue` sample in a new window of **Visual Studio Code.**

 * Install all dependent packages by executing the following command `npm install`.

 * Finally, run the application using the following command `npm run serve`.

 * After the application has started, it will display a URL in the `command line interface`, typically something like (e.g., https://localhost:8080). Copy this URL and paste it into your default web browser.

   ![Dashboard view](https://github.com/boldbi/vue-with-go-sample/assets/129486688/381aa89c-6870-4489-a744-c3617abc7646)

Please refer to the [help documentation](https://help.boldbi.com/embedding-options/embedding-sdk/samples/vuejs-with-go/#how-to-run-the-sample) to know how to run the sample.

## Important notes

In a real-world application, it is recommended not to store passwords and sensitive information in configuration files for security reasons. Instead, you should consider using a secure application, such as Key Vault, to safeguard your credentials.

## Online demos

Look at the Bold BI Embedding sample to live demo [here](https://samples.boldbi.com/embed).

## Documentation

A complete Bold BI Embedding documentation can be found on [Bold BI Embedding Help](https://help.boldbi.com/embedded-bi/javascript-based/).
