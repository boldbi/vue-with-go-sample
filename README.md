# BoldBI Embedding Vue with Go Sample

This Bold BI Vue with Go sample contains the Dashboard embedding sample. In this sample, the Vue application acts as the front-end, and the Go sample act as the back-end application.This sample demonstrates the dashboard rendering with the available dashboard in your Bold BI server.

## Dashboard view

   ![Dashboard view](/images/dashboard.png)

## Prerequisites

* [Go installer](https://go.dev/dl/)
* [Visual Studio Code](https://code.visualstudio.com/download)
* [Node.js](https://nodejs.org/en/)

 > **NOTE:** Node.js v18.17 to v20.15 are supported.

### Supported browsers
  
* Google Chrome, Microsoft Edge, and Mozilla Firefox.

## Configuration

* Please ensure that you have enabled embed authentication on the `embed settings` page. If it is not currently enabled, please refer to the following image or detailed [instructions](https://help.boldbi.com/site-administration/embed-settings/#get-embed-secret-code?utm_source=github&utm_medium=backlinks) to enable it.

    ![Embed Settings](/images/enable-embedsecretkey.png)

* To download the `embedConfig.json` file, please follow this [link](https://help.boldbi.com/site-administration/embed-settings/#get-embed-configuration-file?utm_source=github&utm_medium=backlinks) for reference. Additionally, you can refer to the following image for visual guidance.

    ![Embed Settings Download](/images/download-embedsecretkey.png)
    ![EmbedConfig Properties](/images/embedconfig-file.png)

* Copy the downloaded `embedConfig.json` file and paste it into the designated [location](https://github.com/boldbi/vue-with-go-sample/tree/master/Go) within the application. Please ensure that you have placed it in the application as shown in the following image.

   ![EmbedConfig image](/images/embedconfig-location.png)

## Run a Sample Using Command Line Interface

  1. Open the **command line interface** and navigate to the specified file [location](https://github.com/boldbi/vue-with-go-sample/tree/master/Go) where the project is located.

  2. Run the back-end `Go` sample by using the following command `go run main.go`.

  3. Open the **command line interface** and navigate to the specified file [location](https://github.com/boldbi/vue-with-go-sample/tree/master/Vue) where the project is located.

  4. Install all dependent packages by executing the following command `npm install`.

  5. Finally, run the application using the following command `npm run serve`.

  6. After the application has started, it will display a URL in the `command line interface`, typically something like (e.g., <https://localhost:8080>). Copy this URL and paste it into your default web browser.

## Developer IDE

* [Visual Studio Code](https://code.visualstudio.com/download)

### Run a Sample Using Visual Studio Code

* Open the `Go` sample in **Visual Studio Code**.

* Run the back-end `Go` sample by using the following command in the terminal `go run main.go`.

* Open the `Vue` sample in a new window of **Visual Studio Code**.

* Install all dependent packages by executing the following command `npm install`.

* Finally, run the application using the following command `npm run serve`.

* After the application has started, it will display a URL in the `command line interface`, typically something like (e.g., <https://localhost:8080>). Copy this URL and paste it into your default web browser.

   ![Dashboard view](/images/dashboard.png)

Please refer to the [help documentation](https://help.boldbi.com/embedding-options/embedding-sdk/samples/vuejs-with-go/#how-to-run-the-sample?utm_source=github&utm_medium=backlinks) to know how to run the sample.

## Important notes

In a real-world application, it is recommended not to store passwords and sensitive information in configuration files for security reasons. Instead, you should consider using a secure application, such as Key Vault, to safeguard your credentials.

## Online demos

Look at the Bold BI Embedding sample to live demo [here](https://samples.boldbi.com/embed?utm_source=github&utm_medium=backlinks).

## Documentation

A complete Bold BI Embedding documentation can be found on [Bold BI Embedding Help](https://help.boldbi.com/embedded-bi/javascript-based/?utm_source=github&utm_medium=backlinks).