## RabbitMQ usage example with go, rest api, redis and postgres

![Screenshot from 2022-04-11 14-12-35](https://user-images.githubusercontent.com/61962654/162728318-5d94cda1-50f5-42a0-8df1-b6755dd8b3b3.png)

---

## Running
```shell
docker-compose up
```
```shell
docker-compose down
```

### Access
#### `:8080`: REST API
#### `:5433`: PostgresSQL [`server:M5F3wWtFxkQ8Ra4n`]
#### `:15672`: RabbitMQ Panel [`guest:guest`]


### Tasks sequence:
1. **Parse api `endpoint`** [api.currencylayer.com](http://api.currencylayer.com/live?access_key=b62ffea127d9eeedf1d07c4f6d0439a2)
2. **Publish JSON to `RabbitMQ`**
3. **Send JSON to all `consumers`**
4. **Update rates `Redis` and save them into history table `PostgreSQL`**

## Currencies API
**[currencylayer.com](https://currencylayer.com/)** provides `160` exchange rates

```yaml
API_KEY: b62ffea127d9eeedf1d07c4f6d0439a2 # working
API_EVERY: 60 # get new rates every 60 seconds
```
---
Note: Lots of errors are not wrapped due to project example purpose