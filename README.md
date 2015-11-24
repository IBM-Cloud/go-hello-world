#Go Hello World Overview

This app is a hello world plus a little big of basic functionality using Golang.
The application uses [Gin-Gonic][gin_gonic_url] for the app router and for templating.

Also check out a [blog post](https://developer.ibm.com/bluemix/2015/10/28/getting-started-with-golang-on-bluemix/) on this topic as well for a more in depth look at Go and Gin-Gonic.

## Build Status

[![Build Status](https://travis-ci.org/IBM-Bluemix/go-hello-world.svg?branch=master)](https://travis-ci.org/IBM-Bluemix/go-hello-world) ![Bluemix Deployments](https://deployment-tracker.mybluemix.net/stats/ec4f40db88c2905bab56e7d183b46d53/badge.svg)

## Running the app on Bluemix

You can deploy your own instance of Go Hello World app to Bluemix. To do this, you can either use the _Deploy to Bluemix_ button for an automated deployment or follow the step below to create and deploy your app manually.

[![Deploy to Bluemix](https://deployment-tracker.mybluemix.net/stats/ec4f40db88c2905bab56e7d183b46d53/button.svg)](https://bluemix.net/deploy?repository=https://github.com/IBM-Bluemix/go-hello-world.git)

1. [Install Go][go_install_url].  If you have Go installed skip to step 2.

2. Create a Bluemix Account

    [Sign up for Bluemix][bluemix_signup_url] or use an existing account.

3. Download and install the [Cloud Foundry CLI][cloud_foundry_url] tool


4. `cd` into the `$GOPATH/src` directory

  ```
  cd $GOPATH/src
  ```

5. Clone the app to your local environment from your terminal using the following command:

  ```
  git clone https://github.com/IBM-Bluemix/go-hello-world.git
  ```

6. `cd` into this newly created directory

  ```
  cd go-hello-world
  ```

7. Compile the Go code:

  ```
  $ make
  ```

8. Open the `manifest.yml` file and change the `host` value to something unique

  The host you choose will determinate the subdomain of your application's URL:  `<host>.mybluemix.net`

9. Connect to Bluemix in the command line tool and follow the prompts to log in:

  ```
  $ cf api https://api.ng.bluemix.net
  $ cf login
  ```

10. Deploy your app to Bluemix:

  ```
  $ cf push
  ```


And voila! You now have your very own instance of the Go Hello World app running on Bluemix.

## Run the app locally

1. [Install Go][go_install_url].  If you have Go installed skip to step 2.

2. `cd` into the `$GOPATH/src` directory

  ```
  cd $GOPATH/src
  ```

3. Clone the app to your local environment from your terminal using the following command:

  ```
  git clone https://github.com/IBM-Bluemix/go-hello-world.git
  ```

4. `cd` into this newly created directory

  ```
  cd go-hello-world
  ```


5. Compile the Go code:

  ```
  $ make
  ```

Your app will be automatically be available on port 8080. To access the app, go to localhost:8080 in your browser. Happy developing!

## Contribute
We are more than happy to accept external contributions to this project, be it in the form of issues and pull requests. If you find a bug, please report it via the [Issues section][issues_url] or even better, fork the project and submit a pull request with your fix! Pull requests will be evaulated on an individual basis based on value add to the sample application.

## Troubleshooting

The primary source of debugging information for your Bluemix app is the logs. To see them, run the following command using the Cloud Foundry CLI:

  ```
  $ cf logs go-hello-world --recent
  ```
For more detailed information on troubleshooting your application, see the [Troubleshooting section](https://www.ng.bluemix.net/docs/troubleshoot/tr.html) in the Bluemix documentation.

## Privacy Notice
The go-hello-world sample web application includes code to track deployments to Bluemix and other Cloud Foundry platforms. The following information is sent to a [Deployment Tracker](https://github.com/IBM-Bluemix/cf-deployment-tracker-service) service on each deployment:

* Application Name (application_name)
* Space ID (space_id)
* Application Version (application_version)
* Application URIs (application_uris)

This data is collected from the VCAP_APPLICATION environment variable in IBM Bluemix and other Cloud Foundry platforms. This data is used by IBM to track metrics around deployments of sample applications to IBM Bluemix. Only deployments of sample applications that include code to ping the Deployment Tracker service will be tracked.

### Disabling Deployment Tracking

To disable deployment tracking for this app remove the following 2 lines [line 5](https://github.com/IBM-Bluemix/go-hello-world/blob/master/go-hello-world.go#L5) and [line 13](https://github.com/IBM-Bluemix/go-hello-world/blob/master/go-hello-world.go#L13) in [go-hello-world.go](https://github.com/IBM-Bluemix/go-hello-world/blob/master/go-hello-world.go)

[bluemix_signup_url]: https://ibm.biz/go-hello-world-signup
[cloud_foundry_url]: https://github.com/cloudfoundry/cli/releases
[issues_url]: https://github.com/IBM-Bluemix/go-hello-world/issues
[gin_gonic_url]: https://github.com/gin-gonic/gin
[go_install_url]: https://golang.org/doc/install
