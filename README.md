
- Prometheus -> Responsible for get the metrics
- Grafana -> Dashboard visualization
- cAdivisor -> Monitoring containers

## Grafana
With grafana we can get data from a lot of different service including prometheus;

## How to Run
```bash
docker-compose up -d
```

**url's:** <br />
Prometheus => localhost:9090
cAdivisor => localhost:8080
Grafana => localhost:3000
    user: admin
    pass: admin

### Create a Dashboard
Settings -> Data Source -> Add data source
 1 - Choose Prometheus
 2 - Config
    Name: Prometheus
    URL: http://prometheus:9090 ## container name:port running
 3- Save and Test

Now we are able to create dashboards with prometheus data
We can find a lot of templates created by the community on grafana webpage
https://grafana.com/grafana/dashboards/

cAdivisor dashboard example: https://grafana.com/grafana/dashboards/14282 
 1 - Copy the dashboard id;
 2 - Create -> Import -> Paste the dashboard id
 3 - Use prometheus as the data source and import the dashboard

Create Gauge Dashboard
 1 - Create Dashboard -> Add a new panel
 2 - Config:
    - Time series: Gauge
    - Title: Online users
    - Metrics Browser (Query section): goapp_online_users
    - Threshold: 500-Green, 1000-Yellow, 1500-Red
 3 - Apply

Create Counter Dashboard
 1 - Create Dashboard -> Add a new Panel
 2 - Config
    - Time Series: Stat
    - Metrics Browser: goapp_http_requests_total
    - Title: Total of Requests

Create Histogram Dashboard
 1 - Create Dashboard -> Add a new Panel
 2 - Config
    - Time Series: Graph
    # le = less or equal, +Inf = max value that he can work with
    - Query: goapp_http_request_duration_bucket{le="+Inf"}
    - Legend: {{handler}}
    - Title: Load time x Access


## Alerts
- Alert -> Contact Points
    1 - Add a new contact point
    2 - Choose the channel ex: Telegram, Email, ....
    3 - After set up the cannel we can create an Alert on the respective panel

## Clean up
```
docker-compose down
```
