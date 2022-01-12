# day1 - gorm 增删改查

day1 是 gorm 入门中的入门捏！我还没有学过数据库，Mysql，gorm 都要学，这样才称得上健全（雾）。

## 0 参考

- [Go组件学习——gorm四步带你搞定DB增删改查 - 简书 (jianshu.com)](https://www.jianshu.com/p/1513f55f8192)

- [GORM文档学习总结_wongcony的专栏-CSDN博客_gorm.open](https://blog.csdn.net/wongcony/article/details/79063407)
- [gorm---最全讲解_最爱松露巧克力-CSDN博客_gorm when](https://blog.csdn.net/weixin_36617251/article/details/88849214?ops_request_misc={"request_id"%3A"163928248616780271578525"%2C"scm"%3A"20140713.130102334.."}&request_id=163928248616780271578525&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~top_positive~default-1-88849214.first_rank_v2_pc_rank_v29&utm_term=gorm&spm=1018.2226.3001.4187#模型定义)

## 1 Mysql 和 gorm 认识

Mysql 是一种数据库；

gorm 是 Go 操作数据库的一种框架，Mysql 是其所能操纵的数据库中的一种。

## 2 Mysql 安装和初步设置

### 2.1 下载安装

官网：https://www.mysql.com/downloads/

安装的时候选择全部的部件，之后会用到 workbench。

### 2.2 Mysql 连接库（命令行）

1. 输入 `mysql -u root -p`
2. 输入密码
3. 连接成功

### 2.3 创建数据库（命令行）

1. 输入CREATE DATABASE hitszedu

```shell
mysql> CREATE DATABASE hitszedu;
```

2. 在 workbench 中可以看到这个数据库（关系）。

3. 右键选择 hitszedu 数据库为默认数据库（set as default schema）

   ![Set_as_default_schema][Set_as_default_schema]

### 2.4 workbench SQL 脚本建表

1. 点击上方的 **scripting **的 **new script**

   ![sql脚本建表][sql脚本建表]

2. 输入建表的 sql 语句

例：

```sql
CREATE TABLE `need` (
  `need_id` varchar(128) NOT NULL COMMENT '需求号',
  `need_type` varchar(64) NOT NULL COMMENT '类型:勤工俭学(workStudy)、公益活动(publicBenefit)、特色课程(course)',
  `time` varchar(128) NOT NULL COMMENT '时间（xx月xx日，字符串）',
  `title` varchar(128) NOT NULL COMMENT '需求标题',
  `description` varchar(1024) NOT NULL COMMENT '需求描述',
  `address_name` varchar(1024) NOT NULL COMMENT '地点名（API获取) ',
  `longitude` BIGINT(6) NOT NULL COMMENT '经度',
  `latitude` BIGINT(6) NOT NULL COMMENT '纬度',
  `address_description` varchar(1024) NOT NULL COMMENT '地点描述',
  `reward` varchar(128) NOT NULL COMMENT '报酬',
  `number_of_sign_up` BIGINT(6) NOT NULL COMMENT '可报名人数',
  `publish_time` BIGINT(6) NOT NULL COMMENT '发布时间（时间戳）',
  `publisher_name` varchar(64) NOT NULL COMMENT '发布者名字（如：食堂，学校，xx社团）',
  `publisher_id` varchar(128) NOT NULL COMMENT '发布人（类型为userID）',
  `reviewer_id` varchar(128) NOT NULL COMMENT '审核人（类型为adminID）',
  `state` varchar(64) NOT NULL COMMENT '状态：报名（signUp)(默认)、终止(stop)、待审核（wait)',
  `signed_up` BIGINT(6) NOT NULL COMMENT '已报名人数，便于报名逻辑处理',
  PRIMARY KEY (`need_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '需求';
```

3. 鼠标左键选中，点击闪电符号运行

   ![闪电符号][闪电符号]

4. 建表成功！(记得刷新才能看到）

   ![sql刷新前后][sql刷新前后]

### 2.5 添加数据

先尝试添加一组数据，同样是用 SQL 脚本：

```sql
INSERT INTO `hitszedu`.`need`
(`need_id`, `need_type`, `time`, `title`, `description`, `address_name`, `longitude`, `latitude`, `address_description`, `reward`, `number_of_sign_up`, `publish_time`, `publisher_name`, `publisher_id`, `reviewer_id`, `state`, `signed_up`)
VALUES ('123', '123', '333', '333', '3333', '3333', 3333, 3333, '333333', '3333', '3333', 3333, '3333', '3333', '3333', '3333', 3333);
```

可以通过右键 need 表选择 Select Rows - Limit 1000：

![sql查看数据][sql查看数据]

然后就能查看到数据了。

![sql查看数据2][sql查看数据2]

## 3 gorm 初步使用

```go
package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Need struct {
	NeedID             string `gorm:"type:varchar(128);column:need_id;primary_key"`
	NeedType           string `gorm:"type:varchar(64);column:need_type"`
	Time               string `gorm:"type:varchar(128);column:time"`
	Title              string `gorm:"type:varchar(128);column:title"`
	Description        string `gorm:"type:varchar(1024);column:description"`
	AddressName        string `gorm:"type:varchar(1024);column:address_name"`
	Longitude          int    `gorm:"type:BIGINT(6);column:longitude"`
	Latitude           int    `gorm:"type:BIGINT(6);column:latitude"`
	AddressDescription string `gorm:"type:varchar(1024);column:address_description"`
	Reward             string `gorm:"type:varchar(128);column:reward"`
	NumberOfSignUp     int    `gorm:"type:BIGINT(6);column:number_of_sign_up"`
	PublishTime        int    `gorm:"type:BIGINT(6);column:publish_time"`
	PublisherName      string `gorm:"type:varchar(64);column:publisher_name"`
	PublisherID        string `gorm:"type:varchar(128);column:publisher_id"`
	ReviewerID         string `gorm:"type:varchar(128);column:reviewer_id"`
	State              string `gorm:"type:varchar(64);column:state"`
	SignedUp           int    `gorm:"type:BIGINT(6);column:signed_up"`
}

func main() {

	db, err := gorm.Open("mysql", "root:xxxxxxxx@tcp(127.0.0.1:3306)/hitszedu?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	db.SingularTable(true)
	fmt.Println("连接成功")

	need := &Need{
		NeedID:             "10086",
		NeedType:           "workStudy",
		Time:               "1月2日",
		Title:              "我需要重新集结我的部队",
		Description:        "我的圣堂武士损失惨重",
		AddressName:        "湖北宜昌",
		Longitude:          111111,
		Latitude:           222222,
		AddressDescription: "水电之都",
		Reward:             "100元",
		NumberOfSignUp:     20,
		PublishTime:        1,
		PublisherName:      "lxh",
		PublisherID:        "shirou",
		ReviewerID:         "323232",
		State:              "signUp",
		SignedUp:           6,
	}
	err0 := db.Create(&need).Error
	if err0 != nil {
		panic(err0)
	}
	defer db.Close()
}
```

这里创建了一个 Need 的结构体。然后自己编了一组数据，添加到数据库中。

```go
db, err := gorm.Open("mysql", "root:xxxxxxxxx@tcp(127.0.0.1:3306)/hitszedu?charset=utf8&parseTime=True&loc=Local")
```

这一行的设置需要注意：gorm.Open() 第一个参数表示连接的数据库的类型，第二个参数包括数据库的用户名，密码，Ip地址，端口号，具体数据库和编码规则。

这里我们用的是 Mysql，数据库的用户名是 root，密码是 xxxxxxxx（这个不告诉你啦），tcp 连接（这个复制就好，是本地连接的意思，大概），具体数据库的名字是 hitszedu，编码规则是 utf-8（后面复制就好）

go mod 肯定要用；然后命令行执行 go build 以及运行，或者 VS Code F5 也行。

![gorm操作mysql添加1][gorm操作mysql添加1]

刷新数据库之后可以看到添加成功。

![gorm操作mysql添加2][gorm操作mysql添加2]

## 4 增删改查

增已经有了，就来测试一下其他的吧！(下面的代码并不完整)

### 删

```go
err0 := db.Delete(&need).Error
```

### 改

```go
err0 := db.Model(&need).Update("title", "战局对我们的战士太不利了").Error
```

删除和修改的效果的查看方法和增加是一样的。

### 查(First)

```go
var needResult Need
err0 := db.Where("need_id = ?", "10086").First(&needResult).Error
fmt.Println("result: ", needResult)
```

查询结果：

![gorm查询输出][gorm查询输出]

### 查(Find)

Find 和 First 不同，会查找符合条件的所有数据，所以需要创建一个切片来存储查询的结果。

```go
needResults := make([]Need, 5)
```

 eg1：查找 need_type = workStudy 的数据

```go
needResults := make([]Need, 5)
err0 := db.Where("need_type = ?", "workStudy").Find(&needResults).Error
fmt.Println("result: ", needResults)
if err0 != nil {
	panic(err0)
}
```

查询结果：

```shell
连接成功
result:  [{10086 workStudy 1月2日 我需要重新集结我的部队 我的圣堂武士损失惨重 湖北宜昌 111111 222222 水电之都 100元 20 1 lxh shirou 323232 signUp 6}]
```

eg2：查找所有数据

```go
needResults := make([]Need, 5)
err0 := db.Find(&needResults).Error
fmt.Println("result: ", needResults)
if err0 != nil {
	panic(err0)
}
```

查询结果：

```shell
连接成功
result:  [{10086 workStudy 1月2日 我需要重新集结我的部队 我的圣堂武士损失惨重 湖北宜昌 111111 222222 水电之都 100元 20 1 lxh shirou 323232 signUp 6} {123 123 333 333 3333 3333 3333 3333 333333 3333 3333 3333 3333 3333 3333 3333 3333}]
```

## 5 注意事项

### 5.1 命名问题

注意到 gorm 中结构体的项的命名是驼峰式的；而 workbench SQL 脚本建表中却不是：column 的命名是小写带下划线的，如：

gorm：

```go
NeedType           string `gorm:"type:varchar(64);column:need_type"`
```

SQL：

```sql
`need_type` varchar(64) NOT NULL COMMENT '类型:勤工俭学(workStudy)、公益活动(publicBenefit)、特色课程(course)',
```

如果不这样处理的话是无法用 gorm 操作数据库的，具体原因暂时不明。

### 5.2 错误输出

```go
err0 := db.Create(&need).Error
	if err0 != nil {
		panic(err0)
	}
```

这一部分使得操作错误时有反应，如：

![主键重复panic][主键重复panic]

这里报错是因为数据库中已经有主键 need_id = 10086 的数据了，不能再添加。

<!-- 图片 -->

[Set_as_default_schema]:../_images/Set_as_default_schema.png
[sql脚本建表]:../_images/sql脚本建表.png

[闪电符号]:../_images/闪电符号.png

[sql刷新前后]:../_images/sql刷新前后.png

[sql查看数据]:../_images/sql查看数据.png

[sql查看数据2]:../_images/sql查看数据2.png

[sql查看数据2]:../_images/sql查看数据2.png

[gorm操作mysql添加1]:../_images/gorm操作mysql添加1.png

[gorm操作mysql添加2]:../_images/gorm操作mysql添加2.png

[gorm查询输出]:../_images/gorm查询输出.png
[主键重复panic]:../_images/主键重复panic.png
