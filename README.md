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
