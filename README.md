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
## volt
```php
//值输出
$this->view->variable = 123;
{{ variable }}  == echo 123;
//数组或者对象形式的值
{{ variable.col }}

//判断
{% if variable == 'password' %}
{% elseif variable =='phone' %}
{% endif %}
   
//input-text组件(name属性,k:v,…………)
{{ text_field("email", 'id':'email', "class" : "form-control", "placeholder" : "Email", "required":"required") }}
//提交按钮组件(value按钮值,k:v属性,…………)
{{ submit_button("Sign In", "class":"btn btn-primary btn-block btn-flat") }}

```
