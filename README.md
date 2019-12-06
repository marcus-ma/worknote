# worknote

-------------------------
- [Apcu](#apcu)
- [JB-IDE激活码](#JB-IDE激活码)
- [TODO-source-code](#TODO-source-code)
- [手机端调试工具插件](#手机端调试工具插件)
- [CMD命令](#CMD命令)
- [网页摆脱鼠标快捷键](#网页摆脱鼠标快捷键)
- [常用功能网址](#常用功能网址)
- [在网页中插入网易云播放器](#在网页中插入网易云播放器)
- [win10快速打开powershell](#win10快速打开powershell)
- [ffmpeg常用命令](#ffmpeg常用命令)
- [SQL优化](#SQL优化)
- [MYSQLDump-tips](#MYSQLDump-tips)
- [Phalcon\Model](#Phalcon\Model)
- [PHP技巧](#PHP技巧)
- [PHP封裝的常用函数](#PHP封裝的常用函数)
- [flashSession](#flashSession)
- [volt](#volt)
- [Phalcon\Image\Adapter\GD](#Phalcon\Image\Adapter\GD)
- [searchAction](#searchAction)
- [Bootstrap-table](#Bootstrap-table)
- [MongoDB](#MongoDB)
- [IM常用函数](#IM常用函数)
- [视频流播放](#视频流播放)
- [window-API函数](#window-API函数)
- [itemCF-demo](#itemCF-demo)
- [Trie树关键词搜索](#Trie树关键词搜索)
- [图的知识-Graph](#图的知识-Graph)
- [DAG任务调度器-demo](#DAG任务调度器-demo)
- [MySQL的in的查询结果集按顺序](#MySQL的in的查询结果集按顺序)
- [Go的二三事](#Go的二三事)
-------------------------

## apcu [注意！在php-cli下运行缓存无效的]
```php
//存储K-v-TTL(s)
apcu_store('name', 'marcus', 60);
//读取v,不存在就返回false
var_dump(apcu_fetch('name'));
//删除k
apcu_delete('name')
```
## 常用git命令
1.取消本次修改，状态回到上一次提交的状态：git reset --hard 

## JB-IDE激活码
激活码地址[http://pblog.rzepx.cn/activationcode] </br>

## TODO-source-code
`workflow`
</br>
1:`flow`介绍[https://www.helplib.com/GitHub/article_145952]
</br>
2:`go-workflow`介绍[https://blog.csdn.net/mumushuiding/article/details/90485196] 或者 [https://blog.csdn.net/hotqin888/article/details/78822774] 
</br>
2:`DAG-demo`介绍[https://blog.csdn.net/github_33719169/article/details/84827256]
</br>
`HLS和m3u8`
</br>
1:使用[https://blog.csdn.net/qq_41153478/article/details/84182994]



## 手机端调试工具插件
使用方法教程[https://blog.yurunsoft.com/a/eruda.html] ，感觉比腾讯的vconsole好用</br>


### CMD命令
1:[快速打开网页]start http://xxxx.com
</br>
2:[发起POST请求]
</br>表单:curl http://127.0.0.1:8080/test -X POST -d "name=marcus&token=test" 
</br>json:curl -H "Content-Type:application/json" http://127.0.0.1:8080/test -X POST -d '{"user": "admin", "passwd":"12345678"}'

### 网页摆脱鼠标快捷键
1:[在不同的tag选项卡中切换]ctrl+Tab
</br>
2:[网页中的链接切换选中]Fn+Tab
</br>

### win10快速打开powershell
空白处 + Shift键 + 鼠标右键 

### IDE中的快捷操作
1:[批量删除全部文件]
</br>
先选中第一个开始删除的文件，按住shift键，再选中最后一个要删除的文件，松开shift，会发现全选中了
</br>

### 常用功能网址
1：B站视频在线解析[http://www.xbeibeix.com/bilibili/] </br> [https://pv.vlogdownloader.com/] </br>
2: 网盘搜索 [https://dalipan.com/search?keyword=] </br>
3: 短视频去水印 [https://dy.kukutool.com/] </br>
4: 电子书搜索 [https://www.jiumodiary.com/] </br>

### 在网页中插入网易云播放器
```html
<iframe frameborder="no" border="0" marginwidth="0" marginheight="0" width="330" height="450" src="//music.163.com/outchain/player?type=0&id=12345678&auto=0&height=430"></iframe>
```
上面的id使用的是你想要播放的歌单的id，歌单id就是网易云网页版歌单播放的地址：`https://music.163.com/#/my/m/music/playlist?id=`

### ffmpeg常用命令
1:音视频文件播放倍数改变[将bgm.mp3播放速度改为原来的2倍]：`ffmpeg -i bgm.mp3 -filter_complex "[0:v]setpts=0.5*PTS[v];[0:a]atempo=2.0[a]" -map "[v]" -map "[a]" bgm2.0.mp3`
</br></br>
2:视频文件生成截图[将output.mp4截图第4秒的画面当截图]：`ffmpeg -ss 00:00:03 -y -i output.mp4 -s 500x334 -vframes 1 new.jpg`
</br></br>
3:视频和BGM合并[将output.mp4截图第4秒的画面当截图]：
先去掉视频原声`ffmpeg -i input.mp4 -c:v copy -an no-sound-input.mp4`</br>
再合并bgm`ffmpeg -i no-sound-input.mp4 -i bgm.mp3 -t 7.1 -c y copy output.mp4`
</br>
合并音频和视频，保留视频原声【此时需要将mp3文件放在前面，MP4文件放在后面，否则会没有背景音乐】`ffmpeg -i bgm.mp3 -i input.mp4  -t 7.1 -y output.mp4`
</br></br>
4:推流命令：`ffmpeg -re -stream_loop -1 -fflags +genpts -i input.mp4  -vcodec copy -acodec copy -strict -2 -f flv -y rtmp://localhost:1935/live/room`
</br></br>
5:调节音频的音量大小,混音【增加/减少3dB:+/-3dB】：`ffmpeg -i input.mp3 -af volume=+3dB output.mp3`
</br></br>

## SQL优化
1:【分页查询优化】在分页查询大量数据时候，采用主键索引来定位快,一般的查询采用limit是会逐一统计数据条数。如SELECT * FROM table WHERE id>500000 AND id<=500000+100 的查询速度会比一般分页查询SELECT * FROM table LIMIT 500000,100。
</br></br>
2：https://developer.aliyun.com/article/727673?utm_content=g_1000089764

## MYSQLDump-tips
1:【导出不影响：mysqldump --single-transaction】--single-transaction参数的作用，设置事务的隔离级别为可重复读，即REPEATABLE READ，这样能保证在一个事务中所有相同的查询读取到同样的数据，也就大概保证了在dump期间，如果其他innodb引擎的线程修改了表的数据并提交，对该dump线程的数据并无影响，在这期间不会锁表。

## Phalcon\Model
`initialize` 
```php
public function initialize()
{
       //选择连接的DB
       $this->setConnectionService("DB1");
        //指定数据库，跟上面的一样
        $this->setSchema("DB1");
        //制定连接的表
        $this->setSource("table1");
        //必填，具体机制有待探究
        $this->addBehavior(
            new BlameableLib()
        );
}

```

`validation[校验机制的详细属性参考文章【https://blog.csdn.net/u014691098/article/details/80295304】]`
```php
new Validation\Validator
public function validation()
{
       $validator = new Validation();
       //检验唯一性	
       $validator->add('字段属性', new Uniqueness([
           'message' => '名称已经被使用',
       ]));
       //检验是否存在
       $validator->add('字段属性', new PresenceOf([
           'message' => '名称必需填写',
       ]));
       //自定义检验
       $validator->add('age', new Callback([
          'callback' => function(){
		           return $this->age >= $this->min_age;
          }
             'message' => '年龄一定大于法定最小年龄',
        ]));
       //校验是否存在于集合中
       $validator->add('字段属性', new InclusionIn([
           "domain" => [1,2],
       ]));
-----------------------------------------------------------------------
       //校验是否为URL 【除了提供Url外，还提供Email】
       $validator->add('字段属性url', new Url([
           'message' => '页面地址必需是有效的url',
       ]));
-----------------------------------------------------------------------
        //校验长度
        $validator->add('字段属性phone', new StringLength([
            'min' => '11',
           'max' => '11',
           'messageMaximum' => '手机号不正确：'.$this->phone,
           'messageMinimum' => '手机号不正确：'.$this->phone,
        ]));
        //利用正则校验
        $validator->add('字段属性phone', new Regex([
            'pattern' => '/^1\d{10}$/',
           'message' => '错误的手机号码'.$this->phone,
        ]));
        return $this->validate($validator);
}

```

`静态方法map的用法` 
```php
  //map方法是本系统封装好的一个方法，其作用是查出制定的数据组成关联数组
  //第一个参数传递作为数组key的模型字段，第二个参数传递最为数组value的墨子那个字段
  //第三个参数是sql的where条件
  $kvArray = model::map('id','created_at',['order'=>'created_at asc']);
  //[1=>'2019/2/10',2=>'2019/3/1']
```


`in的用法` 
```php
   $bind['_user'] = $ids;//[1,2,3,4,5,6]
   $records = self::find([
    	'conditions' => 'user_id in ({_user:array})'.$queryWhere,
    	 'bind' => $bind,
   ]);
```

`数据保存失败的提示获取(在控制器层)` 
```php
if (!$model->save()) {
   foreach ($model->getMessages() as $message) {
       $this->flashSession->error($message);
    }
    //准备将之前保存失败的内容显示回在之前的页面上
    $this->tag::setDefaults(get_object_vars($model));
    return $this->dispatcher->forward(['controller' => 'xxx', 'action' => 'xxx',]);
}
```

## PHP技巧
```php
//按照某key来将数组进行排序
$b = [
    ['id'=>1,'na'=>'m'],
    ['id'=>2,'na'=>'a'],
    ['id'=>3,'na'=>'r'],
    ['id'=>4,'na'=>'c']
];
array_multisort(array_column($b, 'id'), SORT_DESC, $b);
var_dump($b);
$b = [
    ['id'=>4,'na'=>'c']
    ['id'=>3,'na'=>'r'],
    ['id'=>2,'na'=>'a'],
    ['id'=>1,'na'=>'m'],
];

//数组合并
$json = ['name'=>'marcus'];
$json+= ['age'=>78];
var_dump($json); //['name'=>'marcus','age'=>78]


//数字格式化
echo number_format(111123435); //111,123,435

```

## PHP封裝的常用函数
```php

//1-友好的显示时间
function friend_date($time)
{
    if (!$time)
        return false;
    $d = time() - intval($time);
    $ld = $time - mktime(0, 0, 0, 0, 0, date('Y')); //得出年
    $md = $time - mktime(0, 0, 0, date('m'), 0, date('Y')); //得出月
    $byd = $time - mktime(0, 0, 0, date('m'), date('d') - 2, date('Y')); //前天
    $yd = $time - mktime(0, 0, 0, date('m'), date('d') - 1, date('Y')); //昨天
    $dd = $time - mktime(0, 0, 0, date('m'), date('d'), date('Y')); //今天
    $td = $time - mktime(0, 0, 0, date('m'), date('d') + 1, date('Y')); //明天
    $atd = $time - mktime(0, 0, 0, date('m'), date('d') + 2, date('Y')); //后天
    if ($d == 0) {
        $fdate = '刚刚';
    } else {
        switch ($d) {
            case $d < $atd:
                $fdate = date('Y年m月d日', $time);
                break;
            case $d < $td:
                $fdate = '后天' . date('H:i', $time);
                break;
            case $d < 0:
                $fdate = '明天' . date('H:i', $time);
                break;
            case $d < 60:
                $fdate = $d . '秒前';
                break;
            case $d < 3600:
                $fdate = floor($d / 60) . '分钟前';
                break;
            case $d < $dd:
                $fdate = floor($d / 3600) . '小时前';
                break;
            case $d < $yd:
                $fdate = '昨天' . date('H:i', $time);
                break;
            case $d < $byd:
                $fdate = '前天' . date('H:i', $time);
                break;
            case $d < $md:
                $fdate = date('m月d日 H:i', $time);
                break;
            case $d < $ld:
                $fdate = date('m月d日', $time);
                break;
            default:
                $fdate = date('Y年m月d日', $time);
                break;
        }
    }
    return $fdate;
}

//demo
var_dump(friend_date(time()));

```


## flashSession
```php
//返回上一页后显示红色提示信息
$this->flashSession->error("primary key was not found");
$this->response->back();
```
![](https://raw.githubusercontent.com/marcus-ma/worknote/master/flashSession-error.png)

```php
//返回上一页后显示绿色提示信息
$this->flashSession->success("ts was created successfully");
$this->response->back();
```
![](https://raw.githubusercontent.com/marcus-ma/worknote/master/flashSession-success.png)

```php
//返回上一页后显示蓝色提示信息
$this->flashSession->notice("ts was created successfully");
//返回上一页后显示黄色提示信息
$this->flashSession->warning("ts was created successfully");
$this->response->back();
```
如果在某action下的volt模板中存在
{{ flashSession.output() }}，
同时该action函数中若又存在$this->flashSession的提示函数，
则每次加载模板都会第一时间把提示框显示出来


## volt
```php
//值输出
$this->view->variable = 123;
{{ variable }}  == echo 123;
//对象形式的值
{{ variable.col }}
//数组形式的值
{{ variable['col']}}
//数组长度
{{ variableArr |length }}

//判断
{% if variable == 'password' %}
{% elseif variable =='phone' %}
{% endif %}
//判断变量是否被定义
{% if profile is defined %}
//结合在action中的$this->tag->setDefault('phone',111),判断是否设置
{% if tag.getValue('phone') %}

 
//遍历
{% for item in set %}
   //对象
   {{ item.name }}
   //数组
   {{ item['name'] }}
{% endfor %}

//kv遍历
{% set numbers = ['one': 1, 'two': 2, 'three': 3], user = 'marcus' %}
{% for name, value in numbers %}
    Name: {{ name }} Value: {{ value }}
{% endfor %}

//遍历时判断
{% for value in numbers if value < 2 %}
    Value: {{ value }}
{% endfor %}


//过滤器【http://docs.iphalcon.cn/reference/volt.html】
{# e or escape filter #}
{{ "<h1>Hello<h1>"|e }}
{# trim filter #}
{{ "   hello   "|trim }}
{# striptags filter #}
{{ "<h1>Hello<h1>"|striptags }}
……………………………………


//注释
{# note: this is a comment
    {% set price = 100; %}
#}


//标签助手
//form表单
{{ form('products/save', 'method': 'post') }}
{{ end_form() }}

//input-text组件(name属性,k:v,…………)
{{ text_field("email", 'id':'email', "class" : "form-control", "placeholder" : "Email", "required":"required") }}
//textarea组件
{{ text_area("content",'id':'content', "class" : "form-control", "placeholder" : "Content",) }}
//提交按钮组件(value按钮值,k:v属性,…………)
{{ submit_button("Sign In", "class":"btn btn-primary btn-block btn-flat") }}
//select组件,id与name都为type，optionValue为[1,2,3,4],k为ov的value，v为ov的页面显示
{{ select("type", [1,2,3,4], 'using': ['id', 'name']) }}
//数字输入框，相当于<input type="number">
{{ numeric_field("id", 'placeholder':'ID', "size" : 5, "class" : "form-control", "id" : "fieldId") }}
//相当于a标签,<a href="ts-campaign/search" class="btn btn-default" >back</a>
{{ link_to("ts-campaign/search", "back", 'class': 'btn btn-default') }}
//隐藏文本框，相当于<input type="hidden" id="id" name="id" value="marcus">
{{ hidden_field("id","value":"marcus") }}
//file文件上传组件,相当于<input type="file" data-name="code_url" id="fieldCode">
{{ file_field('data-name':"code_url", "size" : 30, "class" : "form-control upload-file", "id" : "fieldCode")  }}


//模板继承
//在base.volt中
{# templates/base.volt #}
<!DOCTYPE html>
<html>
    <head>
        {% block head %}
            <link rel="stylesheet" href="style.css" />
        {% endblock %}

        <title>{% block title %}{% endblock %} - My Webpage</title>
    </head>
    <body>
        <div id="content">{% block content %}{% endblock %}</div>
        <div id="footer">
            {% block footer %}&copy; Copyright 2015, All rights reserved.{% endblock %}
        </div>
    </body>
</html>

//在其他组件中
{% extends "templates/base.volt" %}
{% block title %}Index{% endblock %}
{% block head %}<style type="text/css">.important { color: #336699; }</style>{% endblock %}
{% block content %}
    <h1>Index</h1>
    <p class="important">Welcome on my awesome homepage.</p>
{% endblock %}
```



## Phalcon\Image\Adapter\GD
```php
//图像处理【http://docs.iphalcon.cn/reference/volt.html】
//读取图像句柄
$image = new Phalcon\Image\Adapter\GD('C:/Users/53505/Desktop/test.jpg');

//设置生成图片的尺寸
//第一个参数是指定宽度，第二个参数指定高度，第三个参数指定了缩放方式，默认为 PhalconImage::AUTO
$image->resize(200, 200);

//可以添加水印到指定位置
//先获取水印的图片
$mask = new Phalcon\Image\Adapter\GD('watermark.jpg');
//添加到图片的制定位置上
$image->watermark($watermark, -10, -10, 90);

//可以添加文本到指定位置，下面的示例将会添加文本到图像正中间位置：
$image->text('Hello world');

//图像另存`save` 的第一个参数是指定保存路径，如果为空，将覆盖原始文件，第二个参数指定图像质量 [1-100] 之间。
$image->save('unit-tests/assets/production/gd-resize.jpg');

//直接渲染，获取二进制字符串
$bytes = $image->render();
$this->response->setContentType('image/png');
$this->response->setContent($bytes);
$this->response->send();
```

### searchAction
```php
对于所有search行为的控制器，生成表单数据的方法为
if ($this->request->isPost()) {
    $model = new {MODEL}();
    $builder = $model->dataTableBuilder();
    //此处的数组，是用于将的当前的一些模型字段进行自定义的转换
    $transform = [];
    //打个例子：现在在模板的字段有status，department，ip(对应datatable的table中的data-field="xxx")。
    //其中department字段是要与部门model相互关联的
    $transform = [
            'status' => function($data){
	       return model::status[$data->status]??'-';
             },
            'department' => function($data){
	       return $data->departmodel->name;
             },
            'ip' => function($data){
	       //UtiLib::geo方法是本系统封装好的工具函数，会把ip转化成数组,解析为国家、城市、运行商、ip
	       return UtiLib::geo($data->ip)
	       //此时返回的是数组，在模板数据渲染的时候记得要调用{{ partial('widget/extra', ['formatter':['ip']]) }}
             },

    ];
   
    return $this->response->paginate($builder, $transform);
-----------------------------------------------------------------------
//如果想要直接用返回一个对象直接在模板填充，做法是
$transform = [
     'user'' => function($data){
            return isset($data->id) && $data->id ? Users::findFirst($data->id) : null;
      },
];
```
```html
在模板上则为
<th data-field="name" data-formatter="nameFormatter" data-align="center">姓名</th>
<th data-field="sex" data-formatter="sexFormatter" data-align="center">性别</th>
 <th data-field="status" data-formatter="statusFormatter" data-align="center">状态</th>

function nameFormatter(value, row, index){
        return '<span>'+row.user.name+'</span>';
}
function sexFormatter(value, row, index){
        return '<span>'+row.user.sex+'</span>';
}
function statusFormatter(value, row, index){
        return '<span>'+row.user.status+'</span>';
}
```


## Bootstrap-table
详细概念可以参考文章[https://blog.csdn.net/tyrant_800/article/details/50269723] 和 [https://www.cnblogs.com/laowangc/p/8875526.html]

## 游戏g值的生成
```js
let a = [
    "",//app_id
    "",//平台(1:ios,2:安卓)
    "",//app_key
    "",//token_key
    "https://api.xxx.com",//接口地址
    "https://res.xxx.com/main/images/app_id对应的图片.jpg",//要加载的背景图片
    "0",//adid
    "1"//是否为小程序
];
let b = btoa(a.join(","));
let c = b.split("").reverse().join("");
console.log(c)
//打印出来的=符号后面的字符串

```

## MongoDB
```php
//https://segmentfault.com/a/1190000017279755
//连接数据库(client->库->表)"mongodb://username:password@host/database";
$collection = (new MongoDB\Client("mongodb://localhost:27017"))->test->users;
//在本系统的调用方法是：('test' => ['uri' => 'mongodb://127.0.0.1:27017', 'db' => 'test'],)
//$collection = MongoLib::db('test', 'test')->users;

//插入数据
//insertOne插入一条
//insertMany插入多条
$insertOneResult = $collection->insertOne([
    'username' => 'admin',
    'email' => 'admin@example.com',
    'name' => 'Admin User',
]);
//返回插入条数
printf("Inserted %d document(s)\n", $insertOneResult->getInsertedCount());
//返回插入的文档id(字符串)
var_dump($insertOneResult->getInsertedId()->__toString());

//根据条件查询数据条数
$total = $collection->count(['user'=>'marcus']);
var_dump($total);

//查询
$doc = $collection->findOne(['username' => 'admin']);
var_dump($doc->email,$doc->name,$doc->_id->__toString());

//按文档id查询(由于储存的`_id`字段是一个BSON类型的object,所以要按_id字段来查的话,要先进行类型转换）
$mongo_id = new MongoDB\BSON\ObjectId('5ccad4f30af1432e9c002352');
$doc = $collection->findOne(['_id' => $mongo_id]);
//如果没有找到,返回值为null，找到返回值是一个对象,这里我们可以将他强制类型转换`$info = (array)$info`
var_dump((array)$doc);

//多个查询find
//多个查询有两个数组参数,第一个是查询的条件,第二个是选项
//第一个数组的查询条件要注意下`$gt`这是查询选择器,表示大于
//第二个数组的选项意义为:
//`projection`指定返回哪些字段,`skip`跳过0条数据,`limit`查询10条数据,`sort`按age字段正序排序(1为正序,-1为倒序)
$info = $collection->find(
    ['name'=>'Dullcat','age'=>['$gt'=>'1']],
    [
        'projection'=>['id' => 1, 'age' => 1, 'name' => 1],
        'skip'=>0,
        'limit'=>10,
        'sort'=>['age'=>1]
    ]
);
//这时请注意,多个查询返回的是游标(Cursor),需要的进行处理
//您可以用函数转换$info_array = $info->toArray();
//也可以用foreach迭代该对象
foreach($info as $value){
$info_array [] = $value;
}


//更新
//单条更新updateOne，多条更新updateMany
//单个更新和多个更新的用法一模一样,但是要注意的是他的操作符和选项
//update方法有三个数组参数
//第一个数组是查询条件，
//第二数组是更新操作符,例如下面的
//`['$set'=>['age'=>'23']]`,意义为:更改age字段为23
//第三个数组是选项,可以选择各种参数,例如下面的
//`['upsert'=>true]`,意义为当能查询到就进行更改,如果没有查询到就进行新增操作
$result = $collection->updateOne(
    array('username'=>'admin'),
    array('$set'=>array('name'=>'Admin Marcus')),
    array('upsert'=>true)
);
var_dump($result);


//删除
//单个删除deleteOne，多个删除deleteMany
//单个删除和多个删除的用法一模一样,只是单个删除只删除查询到的第一条数据,而多个删除则删除匹配到的所有数据
//单个删除：
$
