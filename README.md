## 1.windows 开发环境准备

### 1.安装 go 1.22.0

> go 版本管理工具：https://github.com/voidint/g

```bash
# 查看本地 go 已安装版本
g ls

# 查看远程可用版本
g ls-remote

# # 查看远程最新稳定版本
g ls-remote stable

g install 1.22.0
g use 1.22.0
```
### 2.安装 docker-desktop

> https://www.docker.com/products/docker-desktop/
> 

### 3.docker 安装并启动 postgres

```bash
# 拉取
docker pull postgres:16-alpine

# 启动 postgres，并设置用户名、密码
docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:16-alpine

# 进入容器，使用 bash （exit 退出）
docker exec -it postgres16 bash
# root 用户登录 psql （\q 退出）
psql -U root

# 上述合并为一行
docker exec -it postgres16 bash -c "psql -U root"
```
使用 `golang` 插件连接数据库，测试连接。或者， 使用 [datagrip](https://www.jetbrains.com/datagrip/) 图形化界面，连接数据库。

### 3.使用[【https://dbdiagram.io/】](https://dbdiagram.io/)创建数据表

<iframe width="560" height="315" src='https://dbdiagram.io/e/665dba85b65d9338795d2fed/665ed607b65d9338797298a5'> </iframe>