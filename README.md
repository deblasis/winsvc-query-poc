# winsvc-query-poc

Simple windows executable used to query for the status of a Windows Service, in this case `Dhcp` (pretty ubiquitous).

The implementation is inspired from https://cs.opensource.google/go/x/sys/+/master:windows/svc/example/

There are two branches:

## [1simple](https://github.com/deblasis/winsvc-query-poc/tree/1simple)

This one uses the basic methods provided by the `x/sys` library that, by default, require the highest permission level: `SC_MANAGER_ALL_ACCESS` ([source](https://cs.opensource.google/go/x/sys/+/master:windows/service.go;drc=6e7872819dc85c039df9bda7e7c628bb19e85d0e;l=17))

Trying to query the Service Manager here fails when connecting when the user running the executable doesn't have Administrator privileges:

![image](https://user-images.githubusercontent.com/29378614/158074761-793232a4-a4ca-4d8e-975f-99605d348f42.png)

😢

## [2leastprivilege](https://github.com/deblasis/winsvc-query-poc/tree/2leastprivilege)

This branch contains calls at a lower level with lower permissions that result in a successful call

![image](https://user-images.githubusercontent.com/29378614/158074772-b71bd2b8-571d-4493-9e36-67a2c5b12e95.png)

😃

# Usage

Switch to one of the above branches and then

```
make run
```
