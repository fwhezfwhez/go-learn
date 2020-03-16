## 使用postgres，数据库增删改查
需要在本机上安装好docker

**拉取镜像，运行,进入容器，进入数据库**

`docker run --rm -itd -p 5432:5432 -e POSTGRES_PASSWORD=123 --name postgres postgres:9.6`

`docker exec -it postgres sh`

`psql -d postgres -U postgres`

**并准备数据库数据**
```sql
-- 建表
create table user_info(
      id serial primary key,
      username varchar not null default '',
      realname varchar not null default '',
      age integer not null default 0,
      raw jsonb,
      created_at timestamp with time zone not null default now()
  );
-- 增加唯一约束
alter table user_info add constraint user_info_unq_uname unique(username);
-- 增
insert into user_info(username,realname, age) values('raoguanghui','饶光辉', 35);
insert into user_info(username,realname,age) values('ft','冯弢', 18);
-- 删
delete from user_info where username='ft';
-- 改
update user_info set age=25 where username='raoguanghui';
-- 查
select * from user_info;

-- 更多数据库操作，参考https://editor.csdn.net/md/?articleId=81738781
```

**运行程序**
```
go get github.com/fwhezfwhez/errorx
go get github.com/jinzhu/gorm
go get github.com/lib/pq
```

`cd ${GOPATH}/src`

`git clone https://github.com/fwhezfwhez/go-learn.git`

`cd ${GOPATH}/src/go-learn`

`cd use-postgres`

`go run main.go`

