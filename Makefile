hello:
	@echo "Hello World"

# 创建 postgres 容器
postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:16-alpine

# 创建 simple_bank 数据库
createDB:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

# 删除 simple_bank 数据库
dropDB:
	docker exec -it postgres16 dropdb simple_bank

# 升级（更新） simple_bank 数据库
migrateUp:
	# use sslmode=disable to fix error: pq: SSL is not enabled on the server
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up

# 降级（回退） simple_bank 数据库
migrateDown:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down


.PHONY: postgres createDB dropDB migrateUp migrateDown