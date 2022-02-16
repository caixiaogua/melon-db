# melondb
### 简单、方便、快速、强大、神奇的数据库引擎

通过http连接数据库，可以在任何语言和框架中轻松调用，通过提交js指令（事务）操作数据库，无需考虑事务锁问题。

```
v4.3 新增简易分布式锁，后文有使用范例
v4.2 新增数据库可视化管理工具：http://127.0.0.1:1688/ui
```

```
db.init()  //数据库初始化（新建或重置数据库）
db.save()  //数据库持久化（数据写入硬盘）
db.load()  //重新加载磁盘数据到内存（忽略未保存数据）
```

```
// 使用范例：

创建数据表：dbc('db.users=[]')
插入一条数据：dbc('db.users.push({id:1,name:"tom",age:25})')
返回一个数据表：dbc('return db.users')
根据条件返回数据：dbc('return db.users.filter(x=>x.id<6)')
根据条件更新数据：dbc('db.users.find(x=>x.id==1).age=29')
删除指定条件的数据，并返回新的数据列表：dbc('db.users=db.users.filter(x=>x.age<36); return db.users')
```

```
// 在go中使用

package main
import (
	"fmt"
	"github.com/caixiaogua/melon-db/melondb"
)
func main() {
	dbc := melondb.Init("test", "http://127.0.0.1:1688/")
	res := dbc(`db.init(); db.arr=[1,2,3,"nice"]; return db;`) //db.init()为数据库初始化，仅需执行一次
	fmt.Println("res", res)
}
```

```
// 在jsgo中使用

function dbc(x){
	var url="http://127.0.0.1:1688/"; //数据服务地址
	var db="test"; //数据库名称，可自定义
	return api.httpPost(url,JSON.stringify({t:db,s:x}),"");
}
function main(){
	var res=dbc("db.init(); db.arr=[1,2,3]; return db"); //db.init()为数据库初始化，仅需执行一次
	return res;
}
```

```
// 在nodejs中可以通过fetch连接数据库

const fetch=require("node-fetch");
async function dbc(s){
	let t="test"; //数据库名称，可自定义
	return fetch("http://127.0.0.1:1688/",{
		method: "POST",
		body: JSON.stringify({t,s})
	}).then(res=>res.text());
}
async function main(){
	// await dbc(`db.init()`); //db.init()为数据库初始化，只需执行一次
	let obj=`{name:"Candy",age:25}`;
	await dbc(`db.user=${obj}`);
	let res=await dbc(`db.count=db.count||0; db.count++; return db.count;`);
	console.log("res", res);
}
main();
```

```
// 在dotnet中使用（以net6的miniapi为例）

var JsonParse = (string x) => System.Text.Json.JsonDocument.Parse(x);
var Stringify = (object x) => System.Text.Json.JsonSerializer.Serialize(x);
var httpPost=(string url, string str)=>{
    using (HttpClient http = new HttpClient())
    {
        var content = new StringContent(str, System.Text.Encoding.UTF8, "application/json");
        HttpResponseMessage res = http.PostAsync(url,content).Result;
        string data = res.Content.ReadAsStringAsync().Result;
        return data;
    }
};
var dbc=(string x)=>{
    var url="http://127.0.0.1:1688/"; //数据服务地址
    var db="test"; //数据库名称，可自定义
    var data=Stringify(new{t=db,s=x});
    return httpPost(url,data);
};
app.MapGet("/melontest", (HttpContext ctx) =>
{
    //向melondb数据库中的arr数据集中添加一条数据
    var newData=Stringify(new{name="Tom",age=21});
    var res=dbc($"db.arr.push({newData});db.save();return db.arr");
    return res;
});
```

```
// v4.3新增功能：简易分布式锁

package main
import (
	"fmt"
	"time"
	"github.com/caixiaogua/melon-db/melondb"
)
var dbc = melondb.Init("test", "http://127.0.0.1:1688/")
func main() {
	// db.getLock(锁名称，请求次数（默认10），请求间隔（默认500ms），锁有效时长（默认5000ms）)
	tk := dbc(`db.getLock("lock")`) //请求分布式锁，成功返回密匙，失败返回空字符
	if tk != "" {                   //获取到锁
		fmt.Println("working......")
		time.Sleep(3 * time.Second) //业务代码
		fmt.Println("work done.")
		dbc(`db.rmLock("lock", "` + tk + `")`) //释放当前锁
	}
}
```
