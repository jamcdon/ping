# Ping
> An application for looking at the closing values of a stock over a set period of days

## Design
- main.go -> entrypoint of the application
- Controllers/ -> Main logic of endpoints
- Models/ -> All structs used in the application, broken down by use case
- Routes -> Application router

## Usage
{applicationURL}/v1/daily/{STOCK_SYMBOL}/{REQUESTED_DAYS_AMOUNT}
- applicationURL:
    - When deployed locally: 0.0.0.0:8080
- STOCK_SYMBOL:
    - Any stock ticker/Symbol
    - Example: MSFT
- REQUESTED_DAYS_AMOUNT:
    - Example: 7

## Deployment
This application lives on [Docker Hub](https://hub.docker.com/repository/docker/jamcdon/ping_app), is built with the Dockerfile, and is designed for deployment using kubernetes from the `ping-app-deployment.yml` manifest.

To deploy you need to run the following commands:
- kubectl create secret generic ping-secrets --from-literal=API_KEY_HERE
- kubectl apply -f ping-app-deployment.yml
- kubectl port-forward service/ping-app-service 8080:80

The application can now be accessed from 0.0.0.0:8080. An example request would be a GET from 0.0.0.0:8080/v1/daily/MSFT/7 . This will show the closing data for the last 7 days from Microsoft's stock.

Example output:

```json
{
    "MetaData": {
        "Symbol": "msft",
        "Days": 7,
        "Closing High": 399.95,
        "Closing Low": 381.87,
        "CloseAvg": 391.5128571428571,
        "Closing Median": 391.79
    },
    "Closing Prices": {
        "2026-03-13": 395.55,
        "2026-03-16": 399.95,
        "2026-03-17": 399.41,
        "2026-03-18": 391.79,
        "2026-03-19": 389.02,
        "2026-03-20": 381.87,
        "2026-03-23": 383
    }
}
```

## Resilience
1. Monitoring
    - Alerting should be configured in your preferred logging provider to look at response codes. The application has some basic logging and status codes that can be used to find increased error rates to notify operations of potential application failure.
    - To further increase alerting awareness, the above metrics can be tied to dashboards and external alerting tools to notify staff of potential immminent failure.
2. Reliability
    - This application is currently configured via kubernetes to have 2 replicas with 1 LoadBalancer. This application currently expects very low usage, has a low importance (no related services attached or dependent) and this is a reasonable configuration.
    - If this were a more important application, more replicas, and replicas of the load balancer will increase reliability by introducing more fallbacks.
    - This application was also created with the intention of living on a single cluster. To increase reliability, adding the service to additional clusters will further introduce more fallbacks, and likely result in increased uptime.
    - While only tangentially related, golang as a language is good at catching many small mistakes made my developers. It is not perfect, can still run into failures such as index out of boun exceptions, or pointer failures. In regards to this application, no pointers were *directly* used. The language choice *should* increase reliability, albeit minorly.
3. Resilience
    - Again, the best way to increase an applications resilience is to add more fallbacks and more deployments. Deploying this application across multiple clouds, multiple zones, or multiple data centers will increase resilience and even protect against natural disasters, and depending on deployment, entire cloud provider failure.
4. Security
    - This application currently has very little safeguards in place from a security engineer's perspective. I would argue due to use case and limited expected usage this is reasonable for the application.
    - To increase security against the static code, API keys have been set as environment variables that are set in the kubernetes deployment. This removes API keys from the code directly, which in turn increases the applications security. The key is currently deployed manually by an engineer, but security could be increased by adding this api key as a secret to the build pipeline of the application.
    - Continuing with API keys, this application would benefit from an API key for usage, whether configured in the application itself, or externally.
    - Injection has low risk in this application. There is no SQL in the code, and due to api usage, is very strict in how parameters are used in the application.
5. Scalability
    - This application has been designed for scalability in mind. There are no direct database connections being made, relies on no dependencies, and is able to process all logic in the application. This application is a great contender for horizontal scaling due to these factors.
    - This application is less likely to need vertical scaling, but depending on conditions, it is able to do so.
    - As this application was built with and for kubernetes deployments, it is exceptionally scaleable across multiple environments, not limited to kubernetes - but could also be scaled using serverless architecture. This application is a great candidate for a serverless deployment, as it has low resource usage, very little "wake" time required, and does not maintain active connections to other services.