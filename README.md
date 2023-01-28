# gok9s
使用 gok9s 命令，配合 k9s 命令可以更方便的切换集群，命名空间。


# 安装和使用
```
go build
./gok9s
```


# 功能
目前有两个功能

- 查询目前已有集群
```
./gok9s query
```

- 切换集群
```
./gok9s switch -c <ClusterName> -n <NamespaceName>

```


# 新功能
后续会增加添加集群等其他的功能。

