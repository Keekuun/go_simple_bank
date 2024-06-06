# Postgresql
## 1.timestamptz vs timestamp
timestamptz 和 timestamp 都是 PostgreSQL 中用于存储日期和时间的数据类型,但它们之间有一些重要的区别:

+ 时区信息:
  - timestamptz (timestamp with time zone) 会存储时区信息,这意味着它可以准确地表示特定时间点在不同时区的对应时间。
  - timestamp (timestamp without time zone) 只存储日期和时间,不包含时区信息。

+ 存储空间:
  - timestamptz 通常需要 8 个字节的存储空间。
  - timestamp 通常需要 4 或 8 个字节的存储空间,取决于精度要求。
  
+ 转换和计算:
  - 使用 timestamptz 时,PostgreSQL 会自动将输入的时间转换为 UTC 时区进行存储和计算。
  - 使用 timestamp 时,PostgreSQL 不会进行任何时区转换,而是直接存储和计算输入的时间。