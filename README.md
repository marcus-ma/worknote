# worknote

## apcu
```php
//存储K-v-TTL(s)
apcu_store('name', 'marcus', 60);
//读取v,不存在就返回false
var_dump(apcu_fetch('name'));
//删除k
apcu_delete('name')
```

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

//判断
{% if variable == 'password' %}
{% elseif variable =='phone' %}
{% endif %}
 
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


## Bootstrap table
详细概念可以参考文章[https://blog.csdn.net/tyrant_800/article/details/50269723] 和 [https://www.cnblogs.com/laowangc/p/8875526.html]

