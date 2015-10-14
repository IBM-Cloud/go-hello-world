#go-hello-world Overview

This app is a hello world plus a little big of basic functionality using Golang.
The application uses [Gin-Gonic][gin-gonic-url] for the app router and for templating.

## Running the app on Bluemix

[![Deploy to Bluemix](https://bluemix.net/deploy/button.png)](https://bluemix.net/deploy)

## Running the app on Bluemix

You can deploy your own instance of Go Hello World app to Bluemix. To do this, you can either use the _Deploy to Bluemix_ button for an automated deployment or follow the step below to create and deploy your app manually.

[![Deploy to Bluemix](https://bluemix.net/deploy/button.png)](https://bluemix.net/deploy)

1. [Install Go][go_install_url].  If you have Go installed skip to step 2.

2. Create a Bluemix Account

    [Sign up for Bluemix][bluemix_signup_url] or use an existing account.

3. Download and install the [Cloud Foundry CLI][cloud_foundry_url] tool


4. Goto the GOPATH directory

  ```
  cd $GOPATH
  ```

5. Clone the app to your local environment from your terminal using the following command:

  ```
  git clone https://github.com/IBM-Bluemix/go-hello-world.git
  ```

6. `cd` into this newly created directory

7. Open the `manifest.yml` file and change the `host` value to something unique

  The host you choose will determinate the subdomain of your application's URL:  `<host>.mybluemix.net`

8. Connect to Bluemix in the command line tool and follow the prompts to log in:

  ```
  $ cf api https://api.ng.bluemix.net
  $ cf login
  ```

9. Deploy your app to Bluemix:

  ```
  $ cf push
  ```
10. Compile the Go code:

  ```
  $ make
  ```


And voila! You now have your very own instance of the Go Hello World app running on Bluemix.

## Run the app locally

1. Create a Bluemix Account

    [Sign up for Bluemix][bluemix_signup_url], or use an existing account.

2. If you have not already, [download node.js][download_node_url] and install it on your local machine.

3. Clone the app to your local environment from your terminal using the following command:

  ```
  git clone https://github.com/IBM-Bluemix/watson-travel-qa.git
  ```

4. cd into this newly created directory

5. Deploy a placeholder app to Bluemix. We need to perform additional steps once it is deployed, so we will add the `--no-start` argument:

  ```
  $ cf push --no-start
  ```

6. Create the Watson Q&A service in Bluemix:

  ```
  $ cf create-service question_and_answer question_and_answer_free_plan travel-qa
  ```

7. Bind the Watson Q&A service to the placeholder app:

  ```
  $ cf bind-service watson-travel-qa travel-qa
  ```

8. Now grab the credentials for this service from the Bluemix dashboard and replace the default configs in `vcap-local.json`

9. Install the required npm and bower packages using the following command:

  ```
  npm install
  ```

10. Start your app locally with the following command:

  ```
  npm start
  ```

Your app will be automatically assigned to a port which will be logged to your terminal. To access the app, go to localhost:PORT in your browser. Happy developing!

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

[bluemix_signup_url]: https://ibm.biz/go-hello-world-signup
[cloud_foundry_url]: https://github.com/cloudfoundry/cli
[issues_url]: https://github.com/IBM-Bluemix/go-hello-world/issues
[gin-gonic-url]: https://github.com/gin-gonic/gin
[go-install-url]: https://golang.org/doc/install
