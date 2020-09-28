# redis-driver

## Install

```
$ go get github.com/innotechdevops/redis-driver
```

## How to use

- Wtih env

```golang
driver := redisdriver.New(redisdriver.ConfigEnv())
```

- With config

```golang
driver := redisdriver.New(redisdriver.Config{
    Host:         os.Getenv("REDIS_HOST"),
    Pass:         os.Getenv("REDIS_PASS"),
    DatabaseName: os.Getenv("REDIS_DATABASE"),
    Port:         redisdriver.DefaultPort,
})
```