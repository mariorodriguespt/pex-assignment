## Performance
The following test were conducted with bombardier on Macbook Pro M2 Max 64GB

### 20 concurrent clients - 2120 requests per second

```shell
➜  bombardier git:(master) ✗ ./bombardier http://127.0.0.1:8080/current -c 20 -l
Bombarding http://127.0.0.1:8080/current for 10s using 20 connection(s)
[==========================================================================================================================================================================================] 10s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec      2120.54     997.64    3753.92
  Latency        9.43ms   104.60ms      2.77s
  Latency Distribution
     50%   288.00us
     75%   341.00us
     90%   435.00us
     95%     1.51ms
     99%   151.83ms
  HTTP codes:
    1xx - 0, 2xx - 21274, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:     1.08MB/s
```

### 100 concurrent clients - 1284 request per second
```shell
➜  bombardier git:(master) ✗ ./bombardier http://127.0.0.1:8080/current -c 100 -l
Bombarding http://127.0.0.1:8080/current for 10s using 100 connection(s)
[==========================================================================================================================================================================================] 10s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec      1284.66    1142.34    3857.30
  Latency       80.81ms   409.84ms      6.04s
  Latency Distribution
     50%   287.00us
     75%   382.00us
     90%     3.91ms
     95%   198.99ms
     99%      2.06s
  HTTP codes:
    1xx - 0, 2xx - 12944, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:   602.48KB/s
```

Increasing the number of concurrent clients decreases the request throughput. Using JWT could be an option, given sqlite
at some point is going to become a bottleneck. I decided not to use JWT since it would give too much power to the client and
the server would lose authority by accepting any input.