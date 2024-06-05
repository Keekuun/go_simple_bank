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

```bash 
# 停止容器 postgres16
docker stop postgres16
# 查看正在运行的容器
docker ps
# 查看所有容器
docker ps -a
# 启动容器 postgres16
docker start postgres16

# 删除容器
docker rm postgres16
```

## 2.数据库相关

### 1.使用[【https://dbdiagram.io/】](https://dbdiagram.io/)创建数据表

<iframe width="560" height="315" src='https://dbdiagram.io/e/665dba85b65d9338795d2fed/665ed607b65d9338797298a5'> </iframe>

### 2.使用[【golang-migrate】](https://github.com/golang-migrate/migrate)管理数据版本及迁移

查看 `migrate cli` 可用命令：

```bash
migrate -help
      goto V       Migrate to version V
      up [N]       Apply all or N up migrations
      down [N] [-all]    Apply all or N down migrations
            Use -all to apply all down migrations
      drop [-f]    Drop everything inside database
            Use -f to bypass confirmation
      force V      Set version V but don't run migration (ignores dirty state)
      version      Print current migration version
```

```bash
migrate create -ext sql -dir db/migration -seq init_schema
```

### 3.使用 [`sqlc`](https://github.com/sqlc-dev/sqlc)

+ 安装 [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)

> 小技巧: windows 系统可以直接将下载的`sqlc.exe` 文件放到 `C:\Users\admin\.g\go\bin` 目录下即可，
> 不需要额外配置环境变量了。（已经配置了golang版本管理工具g）

```bash
$ sqlc version
v1.26.0

$ sqlc help
Usage:
  sqlc [command]

Available Commands:
  compile     Statically check SQL for syntax and type errors
  completion  Generate the autocompletion script for the specified shell
  createdb    Create an ephemeral database
  diff        Compare the generated files to the existing files
  generate    Generate source code from SQL
  help        Help about any command
  init        Create an empty sqlc.yaml settings file
  push        Push the schema, queries, and configuration for this project
  verify      Verify schema, queries, and configuration for this project
  version     Print the sqlc version number
  vet         Vet examines queries

Flags:
  -f, --file string   specify an alternate config file (default: sqlc.yaml)
  -h, --help          help for sqlc
      --no-remote     disable remote execution (default: false)
      --remote        enable remote execution (default: false)

Use "sqlc [command] --help" for more information about a command.
```

+ 初始化： 执行`sqlc init`， 会生成 `sqlc.yaml` 文件
+ 配置 [`postgresql`](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)

> [version 2 版本](# https://docs.sqlc.dev/en/latest/reference/config.html#version-2)

```ymal
version: "2"                            # sqlc 配置版本号                      
cloud:                                  # 云数据库配置相关
  project: "<PROJECT_ID>"
sql:
# postgresql
- schema: "postgresql/schema.sql"       # [用于数据库] 定义的sql所在路径或者文件目录， 比如：./db/migration
  queries: "postgresql/query.sql"       # [用于CRUD] 定义的sql所在路径或者文件目录， 比如：./db/query
  engine: "postgresql"                  # 指定引擎
  codegen:                              # 生成代码相关的配置
  - out: src/authors                    # 输出目录
    plugin: py                          # 插件
    options:
      package: authors
      emit_sync_querier: true
      emit_async_querier: true
      query_parameter_limit: 5
  
  gen:
    go: 
      package: "authors"                # 生成代码的包名，默认取 out 配置的 basename
      out: "postgresql"                 # 输出目录
      sql_package: "pgx/v5"             # pgx/v4, pgx/v5 or database/sql. Defaults to database/sql
      emit_prepared_queries: false
      emit_interface: false
      emit_exact_table_names: false
      emit_json_tags: true              # 生成的structs自动注入json tag
      json_tags_case_style: camel       # 生成的json tag 使用 camel
  database:                             # 数据库配置
    managed: false                      # 云数据库配置相关，默认 false https://docs.sqlc.dev/en/latest/howto/managed-databases.html
    uri: postgresql://${PG_USERNAME}:${PG_PASSWORD}@localhost:5432/${PG_DBNAME}   # 数据库连接DSN
  rules:
    - sqlc/db-prepare
    
# msyql
- schema: "mysql/schema.sql" 
  queries: "mysql/query.sql"
  engine: "mysql"
  gen:
    go:
      package: "authors"
      out: "mysql"
```

## 3.项目相关

### 1.初始化
```bash
go mod init github.com/Keekuun/go_simple_bank

go mod tidy
```