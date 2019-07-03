# worknote

## apcu [注意！在php-cli下运行缓存无效的]
```php
//存储K-v-TTL(s)
apcu_store('name', 'marcus', 60);
//读取v,不存在就返回false
var_dump(apcu_fetch('name'));
//删除k
apcu_delete('name')
```

### CMD命令
1:[快速打开网页]start http://xxxx.com
</br>
2:[发起POST请求]
</br>表单:curl http://127.0.0.1:8080/test -X POST -d "name=marcus&token=test" 
</br>json:curl -H "Content-Type:application/json" http://127.0.0.1:8080/test -X POST -d '{"user": "admin", "passwd":"12345678"}'

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
             }
            'ip' => function($data){
	       //UtiLib::geo方法是本系统封装好的工具函数，会把ip转化成数组,解析为国家、城市、运行商、ip
	       return UtiLib::geo($data->ip)
	       //此时返回的是数组，在模板数据渲染的时候记得要调用{{ partial('widget/extra', ['formatter':['ip']]) }}
             }

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


## Bootstrap table
详细概念可以参考文章[https://blog.csdn.net/tyrant_800/article/details/50269723] 和 [https://www.cnblogs.com/laowangc/p/8875526.html]

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
$result = $collection->deleteOne(array('name'=>'DullCat'));
//多个删除：
$result = $collection->deleteMany(['id' => ['$in' => array[1, 2]]]);
//拿到删除数据的条数：
$count = $result->getDeletedCount();


//去重
$fileName = 'name';
$where = ['id' => ['$lt' => 100]];
$ret = $collection->distinct($fileName,$where);



//聚合
$ops = [
    ['$match' =>['type'=>['$in'=>[2,4]]]],
    //sort顺序不能变，否则会造成排序混乱，注意先排序再分页
    ['$sort' => ['list.create_time' => -1]],
    ['$skip' => 0],
    ['$limit' => 20000],
];
$data = $collection->aggregate($ops);
foreach ($data as $document) {var_dump($document);}


//根据筛选条件查询数据
$filter = [];
////时间范围内
$filter['create_time'] = [
     '$gte' => $start_time,
     '$lte' => $end_time
];
////模糊查询
$filter['name'] = ['$regex' =>$keyword];

$doc = $collection->find($filter, $options)->toArray();
var_dump($doc);




//模糊查询(正则)
$filter = ['username' => ['$in' => [new MongoDB\BSON\Regex('^ad','i')]]];
//OR
$filter = ['username' => ['$regex' => 'ad']];
$doc = $collection->findOne($filter);
var_dump($doc);


-----------------------------------------------------------------------------------------------------|
//demo：采用mongodb的2d平面索引就能完成附近的好友搜索了
$collection = (new MongoDB\Client("mongodb://localhost:27017"))->test->lbs;

//loc是一个经纬度的数组,当然也可以是'loc'=>['lng'=>115.993067,'lat'=> 28.67606]，但官方推荐数组。
$insertOneResult = $collection->insertMany([
    ['name'=>'杨帅哥', 'loc'=>[115.993121,28.676436]],
  ['name'=>'王美眉', 'loc'=>[116.000093,28.679402]],
 ['name'=>'张美眉', 'loc'=>[115.999967,28.679743]],
 ['name'=>'李美眉', 'loc'=>[115.995593,28.681632]],
 ['name'=>'彭美眉',  'loc'=>[115.975543,28.679509]],
 ['name'=>'赵美眉',  'loc'=>[115.968428,28.669368]],
 ['name'=>'廖美眉',  'loc'=>[116.035262,28.677037]],
 ['name'=>'余帅哥',  'loc'=>[116.02477,28.68667]],
 ['name'=>'吴帅哥', 'loc'=>[116.002384,28.683865]],
 ['name'=>'何帅哥',  'loc'=>[116.000821,28.68129]],
]);
printf("Inserted %d document(s)\n", $insertOneResult->getInsertedCount());

//设置2d索引,因为以二维平面上点的方式存储的数据，想要进行LBS查询，那么要设置2d索引
$res=$collection->createIndex(['loc'=>'2d']);
var_dump($res);

//查询附近200米的人
//查询附近的人，首先的指导当前用户所在的经纬度，
//如果不仅想要得到数据还要得到距离，那么可以使用$geoNear指令，
//如果距离自己去计算可以使用$near或者$geoWithin然后在手动计算距离。
//此处采用$geoNear指令查询附近2000m的人。
$ops = [
    ['$geoNear'=>[
        'near'=> [115.999567,28.681813], // 当前坐标
        'spherical'=> true, // 计算球面距离
        'distanceMultiplier'=>6378137, // 地球半径,单位是米,那么的除的记录也是米
        'maxDistance'=>2000/6378137, // 过滤条件2000米内，需要弧度
        'distanceField'=> "distance" // 距离字段别名
    ]]
];
$data = $collection->aggregate($ops);
foreach ($data as $document) {
    var_dump($document->name);
}
//6个人"何帅哥" "张美眉""王美眉" "吴帅哥" "李美眉" "杨帅哥"
-----------------------------------------------------------------------------------------------------|

```

## IM常用函数 详细概念可以参考文章[https://blog.csdn.net/tiantangyouzui/article/details/77540611/] 
```html
<div>
    <!--调用相机-->
    <input type="file" id="camer" accept="image/*" capture="camera" onchange="takingPhoto(this)">
    <!--相片预览区域-->
    <img id="ig" src="">
    <!--相片发送-->
    <button id="img-send" onclick="upload()">发送图片</button>
</div>

<div>
    <!--录音-->
    <input value="请按住说话" type="button" id="speak" style="margin-top: 5px;width: 100%; text-align: center"  />
    <audio id="audio" style="display: none"></audio>
    <audio id="audio4play"></audio>
</div>
```

```js
 function takingPhoto(obj) {
        photoPreview(obj,document.getElementById("ig"))
    }
    
    function upload(){
        uploadfile("upload",document.getElementById("camer"),function (res) {
            alert(res);
        })
    }
    //照片预览
    function photoPreview(fileDom,imgDom) {
        var file = fileDom.files[0],
            reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = function(e){
            imgDom.setAttribute("src", e.target.result);
        };
    }
    //文件上传
    function uploadfile(url,fileDom,callback){
        //H5新特性
        let formdata = new FormData();
        //获得一个文件file.files[0]
        formdata.append("file",fileDom.files[0]);
        //console.log(formdata);
        // formdata.append("filetype",".png");//.mp3制定后缀
        var xhr = new XMLHttpRequest();//ajax初始化
        var url = "http://"+location.host+"/"+url;
        xhr.open("POST",url,true);
        xhr.onreadystatechange = function () {
            if(xhr.readyState == 4&& xhr.status == 200){
                callback(JSON.parse(xhr.responseText))
            }
        };
        xhr.send(formdata);
    }
    
    
    //按钮按住触发事件
     document.getElementById("speak").addEventListener('touchstart', startrecorder, false);
     //按钮松开触发事件
     document.getElementById("speak").addEventListener('touchend', stoprecorder, false);
     //录音函数
     function startrecorder(){
        let audioTarget = document.getElementById('audio');
        var types = ["video/webm",
            "audio/webm",
            "video/webm\;codecs=vp8",
            "video/webm\;codecs=daala",
            "video/webm\;codecs=h264",
            "audio/webm\;codecs=opus",
            "video/mpeg"];
        var suporttype ="";
        for (var i in types) {
            if(MediaRecorder.isTypeSupported(types[i])){
                suporttype = types[i];
            }
        }
        if(!suporttype){
            alert("编码不支持");
            return ;
        }

        this.duration = new Date().getTime();
        navigator.mediaDevices.getUserMedia({audio: true, video: false})
            .then(function(stream){
                this.showprocess = true;
                this.recorder = new MediaRecorder(stream);
                audioTarget.srcObject = stream;
		//配置数据可用的回调,发生在按钮松开的时候	
                this.recorder.ondataavailable = (event) => {
                    console.log("ondataavailable");
                    uploadblob("upload",event.data,".mp3",res=>{
		        //记录语音长度
                        var duration = Math.ceil((new Date().getTime()-this.duration)/1000);
			//拼接语音的信息
                        this.sendaudiomsg(res.data,duration);
                    });
                    stream.getTracks().forEach(function (track) {
                        track.stop();
                    });
                    this.showprocess = false
                };
                this.recorder.start();
            }.bind(this)).
        catch(function(err){
            console.log(err);
            alert(err);
            this.showprocess = false
        }.bind(this));
    }

    function stoprecorder() {
        if(typeof this.recorder.stop=="function"){
            this.recorder.stop();
        }
        this.showprocess = false;
        console.log("stoprecorder")

    }

    function uploadblob(uri,blob,filetype,fn){
        var xhr = new XMLHttpRequest();
        xhr.open("POST","http://"+location.host+"/"+uri, true);
        // 添加http头，发送信息至服务器时内容编码类型
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 304)) {
                fn.call(this, JSON.parse(xhr.responseText));
            }
        };
        var _data=[];
        var formdata = new FormData();
        formdata.append("filetype",filetype);
        formdata.append("file",blob);
        xhr.send(formdata);
    }
    
    //图片压缩
    function compress() {
    //demo::https://blog.csdn.net/yasha97/article/details/83629510
    	//<div>
    	   //<img id="img" src="">
           //<input id="file" type="file" onchange="compress()">
	//</div>
        let fileObj = document.getElementById('file').files[0]; //上传文件的对象
        let reader = new FileReader();
        reader.readAsDataURL(fileObj);
        reader.onload = function(e) {
            let image = new Image(); //新建一个img标签（还没嵌入DOM节点)
            image.src = e.target.result;//转成base格式了
            image.onload = function() {
                let canvas = document.createElement('canvas'),
                    context = canvas.getContext('2d'),
                    imageWidth = image.width / 2,    //压缩后图片的大小
                    imageHeight = image.height / 2,
                    data = '';

                canvas.width = imageWidth;
                canvas.height = imageHeight;

                context.drawImage(image, 0, 0, imageWidth, imageHeight);
                data = canvas.toDataURL('image/jpeg');

                //压缩完成
                document.getElementById('img').src = data;
            }
        }
    }
    
    //图片粘贴事件
    document.addEventListener('paste', function (e) {
        console.info(e);
        const cbd = e.clipboardData;
        const fr = new FileReader();
        const html = '';
        for (let i = 0; i < cbd.items.length; i++) {
            const item = cbd.items[i];
            console.info(item);
            if (item.kind == "file") {
                const blob = item.getAsFile();
                if (blob.size === 0) {
                    return;
                }
                //此处的blob就是为图片的blob对象，要上传的话就直接传它即可
                console.info(blob);
                fr.readAsDataURL(blob);
                fr.onload = function (e) {
		    //显示文件
		    document.getElementById("ig").src = this.result;
                }
            }
        }
    });
    
    
     
```

## 视频流播放(采用blob对象)

```html
<div>
    <video src=""></video>
</div>
```

```js

    //资源链接
    var assetURL = '//f.us.sinaimg.cn/003bcPotlx07uDCIGR9m01041200nK1u0E010.mp4?label=mp4_720p&template=1280x720.23.0&Expires=1560485362&ssig=rIbEZ81VFV&KID=unistore,video';
    //video标签
    var video = document.querySelector('video');
    //创建MediaSource对象
    var mediaSource = new MediaSource;
    //相互关联
    video.src = URL.createObjectURL(mediaSource);
    //添加注册事件触发当上述“连接”结束之后的回调处理。注意：回调处理就是需要赋值 视频数据 的地方
    mediaSource.addEventListener('sourceopen', sourceOpen);

    function sourceOpen(){
        var mediaSource = this;
        //构建一个存放视屏数据的 Buffer
        var sourceBuffer = mediaSource.addSourceBuffer('video/mp4; codecs="avc1.42E01E, mp4a.40.2"');

        fetchAB(assetURL,function (buf) {
            //在往 Buffer 中存放数据结束之后会触发事件
            sourceBuffer.addEventListener('updateend', function () {
                //通过注册这个事件的回调，可以知晓数据已经加载完毕
                mediaSource.endOfStream();
                //视频开始播放
                video.play();
            });
            sourceBuffer.appendBuffer(buf); // buf is the arraybuffer to store the video data
        });
    }

    //获取视频的数据源
    function fetchAB (url, cb) {
        console.log(url);
        var xhr = new XMLHttpRequest;
        xhr.open('get', url);
        xhr.responseType = 'arraybuffer';
        xhr.onload = function () {
            cb(xhr.response);
        };
        xhr.send();
    }

```

## js分享第三方平台
### 分享到微博
```js
    //title微博内容,url添加到微博内容的链接,picurl图片的url
    function shareToWeibo(title,url,picurl){
            var sharesinastring='http://v.t.sina.com.cn/share/share.php?title='+title+'&url='+url+'&content=utf-8&sourceUrl='+url+'&pic='+picurl;
            window.open(sharesinastring,'newwindow','height=400,width=400,top=100,left=100');
    }
    //微博分享API参数说明
    //&url=将页面地址转成短域名，并显示在内容文字后面。(可选，允许为空)【在手机端上点击"返回继续浏览"会跳到该链接上】
    //&title=分享时所示的文字内容，为空则自动抓取分享页面的title值(可选，允许为空)
    //&pic=自定义图片地址，作为微博配图(可选，允许为空)
    
     //由于window.open很容易被浏览器拦截，可以直接这样:
    //<a href="http://service.weibo.com/share/share.php?appkey=&title=&url=&pic=&searchPic=false&style=simple" target="_blank">
    //当然，如果你采用这种方式了，记得一定要把appkey=&title=&url=&pic=里你自定义的参数做一次encodeURIComponent之后再嵌入。
    
```




## MySQL in 查询，并通过 FIELD 函数按照查询条件顺序返回结果
详细概念可以参考文章[http://martin91.github.io/blog/articles/2015/09/13/mysql-in-query-and-order-by-field-function/] 

## 函数柯里化
```js
//原本
    function bling(a,other) {
        return a+" "+other;
    }
    //每次都要传入基组件"a"
    let bCss = bling("a","b");
    let cCss = bling("a","c");

//柯里化
    function bling(a){
        return (other)=>{
            return a+" "+other;
        }
    }
    //基组件可以自由设置
    let aCss = bling("a");
    //再可以自由搭配
    let bCss = aCss("b");
    let cCss = aCss("c");
```

## Go的二三事
### defer执行顺序
多个defer的执行顺序是以栈的特性先进后出来执行，越写在前面就越后执行


### 没有前置运算符
不存在++a或者--a，只有a++和a--

### 统计汉字字符串个数
```go
import (
   "fmt"
)
func main(){
     str := "中文"
     fmt.Println(len(str))//6
     c:=[]rune(str)
     fmt.Println(len(c))//2
     
     //golang中string底层是通过byte数组实现的。
     //中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
     //因此直接使用len(str)得出6
     //而是经过rune转换后就会得出想要的结果
     //byte 等同于int8，常用来处理ascii字符
     //rune 等同于int32,常用来处理unicode或utf-8字符
}

```


### 获取命令行参数
```go
//文件名 test.go
import (
   "os"
   "fmt"
)

func main(){
   //获取
   fmt.PrintLn(os.Args)
   if len(os.Args)>1{
   	//os.Args[0]为文件所在详细路径
   	fmt.PrintLn("参数"，os.Args[1])
   }
   //退出
   os.Exit(0)
}

> go run test.go chao
chao
```

### 快速设置连续值
```go
const (
  Monday = iota + 1
  Tuesday 
  Wednesday
  Thursday
  Friday
  Saturday
  Sunday
)

const (
  Readable = 1 << iota
  Writable 
  Executable
)
func main(){
   a := 7 //0111
   fmt.PrintLn(a&Readable==Readable,a&Writable==Writable,a&Executable==Executable)
   //true true true
}

```


### string类型初始化为空字符串
```go
func main(){
   var s string
   fmt.PrintLn("*"+s+"*")//**
   fmt.PrintLn(len(s))//0
   if s == ""{
   	fmt.PrintLn("hello")
   }
}
```



### 使用runtime.Gosched()来让协程交出控制权[I/O操作(如fmt.Println)或者select会自动进行控制权切换]
```go
import (
   "runtime"
   "time"
   "fmt"
)

func main(){
   var a [10]int
   for i:=0;i<10;i++{
   	go func(i int){
		for{
		   //指令操作,不会自动交出控制权
		   a[i]++
		   //让出控制权
		   runtime.Gosched()
		}
	}(i)
   }
   time.Sleep(time.Millisecond)
   fmt.Println(a)
}
```

### 不可思议的接口函数
#### 利用String()接口使结构体可被直接Println
```go
import "fmt"

type Student struct {
   Content string
}

//String接口,当调用者直接fmt.Println结构体时候会调用该方法
func (s *Student) String() string {
	return fmt.Sprintf(
		"Student:{Contents=%s}",s.Content)
}

func main(){
   s:=&Student{Content:"I am student"}
   fmt.Println(s)//Student:{Contents=I am student}
}
```


### 计算函数耗时的函数式编程
```go
import (
   "fmt"
   "time"
)
//计算传入函数执行的耗时
func timeSpent(innerFunc func(op int)int) func(op int) int {
	return func(i int) int {
		//记录开始时间
		start := time.Now()
		ret := innerFunc(i)
		//计算出整个耗时
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func main(){
   test:= func(op int) int {
	 time.Sleep(time.Second)
	 return op
   }

   ret := timeSpent(test)
   ret(30)
}


```


### 使用bufio来提升文件写入速率
```go
import (
	"bufio"
	"fmt"
	"os"
)
func printFile(filename string)  {
	file,err:= os.Create(filename)
	if err!=nil {
		panic(err.Error())
	}
	defer file.Close()

	//使用buffer先将内容写入其中
	writer := bufio.NewWriter(file)
	//最后再一起把内容刷到文件中去
	defer writer.Flush()

	//将内容写到buffer中
	fmt.Fprintln(writer,"hello")
}
```

### 使用bufio和io.Reader来构建万能读取函数
```go
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)
func printFileContents(reader io.Reader)  {
	//创建buffer读取器，将传入的io型reader
	scanner := bufio.NewScanner(reader)
	//逐行扫描打印
	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
}

func main(){
   //demo1：读取文件的内容
   file,err:= os.Open(filename)
   if err!=nil {
	   panic(err.Error())
   }
   defer file.Close()
   printFileContents(file)
   
   //demo2：读取字符串
   //通过``来构建多行的内容字符串
   s1 := `abc"d"
	kkk
	1234
	p`
   printFileContents(strings.NewReader(s1))
}

```

### 服务器统一出错处理(把error都传到中间件来统一处理)
```go
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error
//定义展示给用户看的信息接口
type userError interface {
	error
	Message() string
}
//中间件
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		var (
			err  error
			code int
		) 
		//捕捉不可估计的错误
		defer func() {
			if rec = recover();rec!=nil{
				code = http.StatusInternalServerError
				log.Printf("Panic: %v",rec)
				http.Error(w,http.StatusText(code),code)
			}
		}()
		if err = handler(w,r);err!=nil{
			//返回定义好的友好的给用户看的错误信息
			if userErr,ok:=err.(userError);ok{
				http.Error(w,userErr.Message(),http.StatusBadRequest)
				return
			}
			code = http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w,http.StatusText(code),code)
		}
		
	}
}


//实现展示给用户看的信息接口
type userErrorShow string
func (e userErrorShow) Error() string {
	return e.Message()
}
func (e userErrorShow) Message() string {
	return string(e)
}
//把可能出现的错误都返回出去，把错误交给errWrapper中间件处理
func HandleFileList(w http.ResponseWriter, r *http.Request) error {
	var (
		path string
		err error
		file *os.File
		all []byte
		prefix string
	)
	prefix = "/list/"
	//当用户在浏览器没有按照指定前缀输入，就返回指定的用户看的错误信息
	if strings.Index(r.URL.Path,prefix)!=0 {
		return userErrorShow("path must start with "+prefix)
	}
	//获取后缀的文件名
	path = r.URL.Path[len(prefix):]
	if file,err = os.Open(path);err!=nil{
		return err
	}
	defer file.Close()
	//读取文件
	if all,err = ioutil.ReadAll(file);err!=nil{
		return err
	}

	w.Write(all)
	return nil
}


func main(){
   //读取文件夹中的文件并展示出来
   http.HandleFunc("/list/",errWrapper(HandleFileList))
   http.ListenAndServe(":8888",nil)
}
```






