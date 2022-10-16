# melondb
### 简单、方便、高效的数据库引擎

通过提交js指令（事务）操作数据库，无需考虑事务锁问题。

#### 欢迎加入QQ群：739721147

#### v6.0更新：使用更高效的压缩加密格式存储数据（“.dbx”），存储大小减少90%

##### 启动服务后，可通过 http://localhost:1688/ui 访问webUI

```
// 使用范例：

创建数据表：dbc('db.users=[]')
插入一条数据：dbc('db.users.push({id:1,name:"tom",age:25})')
返回一个数据表：dbc('return db.users')
根据条件返回数据：dbc('return db.users.filter(x=>x.id<6)')
根据条件更新数据：dbc('db.users.find(x=>x.id==1).age=29')
删除指定条件的数据，并返回新的数据列表：dbc('db.users=db.users.filter(x=>x.age<36); return db.users')

系统默认为内存数据库，如果需要持久化数据（写入硬盘），可使用 db.Save() 命令，例如：
dbc('db.users.push({name:'Jerry",age:19});db.Save();')

使用 db.Backup() 命令可将当前数据库文件备份到 backup 文件夹。
dbc('return db.Backup()')	//备份成功返回true，失败返回false

dbc('return db.Export("xxx.json")')	//将当前数据库导出为json文件
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
// 在jsgo中使用

let dbc=x=>api.httpPost("http://127.0.0.1:1688/test",encodeURIComponent(x)); //test为数据库名
function main(){
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
