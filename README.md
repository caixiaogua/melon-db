# melondb
### 简单、方便、高效、灵活的数据库引擎（JSQL）

通过提交js指令（事务）操作数据库，无需考虑事务锁问题。

#### 欢迎加入QQ群：739721147

#### v6.2更新：
1. 优化程序逻辑，数据库并发性能提升20%
2. 精简优化，程序体积减小30%，并提供更友好的js错误提示
3. 使用更高效的压缩加密格式存储数据（“.dbx”），存储大小减少90%
4. 新增db.AutoID()和db.FormatDate()方法，可以获取自动id和格式化日期时间

##### 启动服务后，可通过 http://localhost:1688/ui 访问WebUI数据管理工具

```
// 使用范例：

创建数据表：dbc('db.users=[]')
插入一条数据：dbc('db.users.push({id:1,name:"tom",age:25})')
返回一个数据表：dbc('return db.users')
根据条件返回数据：dbc('return db.users.filter(x=>x.id<6)')
根据条件更新一条数据：dbc('db.users.find(x=>x.id==1).age=29')
根据条件更新多条数据：dbc('db.users.filter(x=>x.id<10).forEach(x=>x.age=29)')
删除指定条件的数据，并返回新的数据列表：dbc('db.users=db.users.filter(x=>x.age<36); return db.users')

添加数据时使用自动id：dbc('db.users.push({id:db.AutoID(db.users),name:"tom",age:25})')
db.AutoID(arr, key='id')带参数时获取arr的自增id，不带参数则获取唯一字符串id如：'luzj0gq6-4izv'

db.FormatDate(t)带参数时获取时间戳t对应的格式化日期，不带参数则当前日期时间，如：2024-04-14 20:22:49

系统默认为内存数据库，如果需要持久化数据（写入硬盘），可使用 db.Save() 命令，例如：
dbc('db.users.push({name:'Jerry",age:19});db.Save();')

使用 db.Backup() 命令可将当前数据库文件备份到 backup 文件夹。
dbc('return db.Backup()')	//备份成功返回true，失败返回false

使用 db.Export(filename,data?) 命令将当前数据库（或指定数据）导出为json文件
dbc('return db.Export("xxx.json")')	//将当前数据库导出为xxx.json文件
dbc('return db.Export("users.json", db.users)')	//将db.users导出为users.json文件

复杂事务，也无需考虑锁问题
例：user1向user2转账100元，金额不足则返回错误信息及实际余额
dbc(`
  let user1=db.users.find(u=>u.uname=="user1");
  let user2=db.users.find(u=>u.uname=="user2");
  if(user1.money>=100){
    user1.money-=100;
    user2.money+=100;
    return {ok:1};
  }else{
    return {error:"余额不足",money:user1.money};
  }
`);

另外，还可以将以上逻辑代码注册为云函数，例如
//执行以下代码（仅需一次），在数据库端创建 db.fn_pay 函数
dbc(`
	db.fn_pay=function(from,to,num){
	  let user1=db.users.find(u=>u.uname==from);
	  let user2=db.users.find(u=>u.uname==to);
	  if(user1.money>=num){
	    user1.money-=num;
	    user2.money+=num;
	    return {ok:1};
	  }else{
	    return {error:"余额不足",money:user1.money};
	  }
	}
`);

//之后，可在任意地方调用（跨线程，跨进程，跨服务，跨系统，跨平台）
dbc(`db.fn_pay("user1","user2",99)`);
```


```
// 在go中使用

package main
import (
	"fmt"
	"github.com/caixiaogua/melon-db/melondb"
)
func main() {
	dbc := melondb.Init("http://127.0.0.1:1688/test")
	res := dbc(`db.arr=[1,2,3,"nice"]; return db;`)
	str := string(res)
	fmt.Println("res", str)
}
```
#### 了解jsgo：https://github.com/caixiaogua/jsgo
```
// 在jsgo7.0以下版本中使用
let dbc=x=>api.httpPost("http://127.0.0.1:1688/test",encodeURIComponent(x)); //test为数据库名
function main(ctx){
	let res=dbc("db.arr=[1,2,3]; return db");
	return res;
}

// 在jsgo7.0及以上版本中使用
let dbc=api.melondb("http://192.168.1.200:1688/test"); //test为数据库名
function main(ctx){
	let res=dbc("db.arr=[1,2,3]; return db");
	return res;
}
```

```
// 在nodejs中可以通过fetch连接数据库

import fetch from "node-fetch";
async function dbc(s){
	return fetch("http://127.0.0.1:1688/test",{
		method: "POST",
		body: encodeURIComponent(s)
	}).then(res=>res.text());
}
async function main(){
	let res=await dbc(`db.count=db.count||0; db.count++; return db.count;`);
	console.log("res", res);
}
main();
```

##### 其它语言可以参考以上范例自己构建post请求
