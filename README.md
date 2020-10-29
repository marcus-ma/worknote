# worknote

-------------------------
- [JB-IDE激活码](#JB-IDE激活码)
- [Github的CDN](#Github的CDN)
- [TODO-source-code](#TODO-source-code)
- [数据分析思维](#数据分析思维)
- [手机端调试工具插件](#手机端调试工具插件)
- [CMD命令](#CMD命令)
- [网页摆脱鼠标快捷键](#网页摆脱鼠标快捷键)
- [常用功能网址](#常用功能网址)
- [在网页中插入网易云播放器](#在网页中插入网易云播放器)
- [win10快速打开powershell](#win10快速打开powershell)
- [ffmpeg常用命令](#ffmpeg常用命令)
- [SQL优化](#SQL优化)
- [计算请求平均耗时百分比](#计算请求平均耗时百分比)
- [Golang环境安装下载](#Golang环境安装下载)
- [Golang调用系统默认浏览器打开指定链接](#Golang调用系统默认浏览器打开指定链接)
- [go.mod的使用，告别GOPATH](#go.mod的使用，告别GOPATH)
- [GO的常用代码段](#GO的常用代码段)
- [MYSQLDump-tips](#MYSQLDump-tips)
- [利用差分数组和前缀和来统计每个位置的出现次数](#利用差分数组和前缀和来统计每个位置的出现次数)
- [xmSelect常用法](#xmSelect常用法)
- [PHP技巧](#PHP技巧)
- [PHP封裝的常用函数](#PHP封裝的常用函数)
- [Bootstrap-table](#Bootstrap-table)
- [MongoDB](#MongoDB)
- [js小功能](#js小功能)
- [IM常用函数](#IM常用函数)
- [视频流播放](#视频流播放)
- [window-API函数](#window-API函数)
- [itemCF-demo](#itemCF-demo)
- [Trie树关键词搜索](#Trie树关键词搜索)
- [图的知识-Graph](#图的知识-Graph)
- [DAG任务调度器-demo](#DAG任务调度器-demo)
- [获取数据流中的中位数](#获取数据流中的中位数)
- [GeoHash算法在LBS的应用](#GeoHash算法在LBS的应用)
- [MySQL的in的查询结果集按顺序](#MySQL的in的查询结果集按顺序)
- [Go的二三事](#Go的二三事)
-------------------------

## 常用git命令
1.取消本次修改，状态回到上一次提交的状态：git reset --hard </br>
2.提交时忽略指定文件：git checkout 文件路径 `首先先git status 查询出所有需要提交的文件然后将需要忽略的文件路径复制下来，执行git checkout 文件路径，然后git status 发现被忽略的文件已经不再其中，然后进行提交` </br>
3.克隆远程指定分支：git clone -b dev http://xxxxx.git </br>
4.拉取别人的分支到本地 ：git fetch origin dev </br>
5.查看拉取代码的方式：git config --local -e </br>
6.遇到http拉取禁止解决办法(window)：[https://blog.csdn.net/qq_34817440/article/details/106432747] </br>

## JB-IDE激活码
激活码地址[http://pblog.rzepx.cn/activationcode] </br>
[http://idea.medeming.com/jets/]</br>

## Github的CDN
例子(获取marcus-ma/SSE-Event的SSE.js)：[https://cdn.jsdelivr.net/gh/marcus-ma/SSE-Event@master/SSE.js]


## TODO-source-code
`协同推荐SlopeOne 算法`
</br>
1:[https://www.cnblogs.com/huangxincheng/archive/2012/11/22/2782647.html]
</br>
2:[
https://blog.csdn.net/xidianliutingting/article/details/51916578]
</br></br>

`前端遇上Go: 静态资源增量更新的新实践 `
</br>
1:[https://www.sohu.com/a/241080725_661013]
</br></br>

`用户行为分析系列`
</br>
1:`需求场景`[https://blog.csdn.net/qq_28773159/article/details/79752718#10006-weixin-1-52626-6b3bffd01fdde4900130bc5a2751b6d1]
</br>
2:`数据的源头`[https://blog.csdn.net/qq_28773159/article/details/79778061]
</br>
3:`A/B 测试`[https://blog.csdn.net/qq_28773159/article/details/79923749]
</br>
4:`热力图`[https://blog.csdn.net/qq_28773159/article/details/79933652]
</br>
5:`渠道推广`[https://blog.csdn.net/qq_28773159/article/details/79939164]
</br></br>


`crc32在做一致性哈希以及哈希环中遇到的坑(对标PHP和golang) `
</br>
1:[https://my.oschina.net/u/1186749/blog/795102]
</br></br>

`每日生产万亿消息数据入库，腾讯如何突破大数据分析架构瓶颈(Bitmap应用) `
</br>
1:[https://www.infoq.cn/article/BDizxoQ766UKblld-E1u]
</br></br>

`workflow`
</br>
1:`flow`介绍[https://www.helplib.com/GitHub/article_145952]
</br>
2:`go-workflow`介绍[https://blog.csdn.net/mumushuiding/article/details/90485196] 或者 [https://blog.csdn.net/hotqin888/article/details/78822774] 
</br>
2:`DAG-demo`介绍[https://blog.csdn.net/github_33719169/article/details/84827256]
</br></br>
`HLS和m3u8`
</br>
1:使用[https://blog.csdn.net/qq_41153478/article/details/84182994]
</br></br>
`golang热重启`
</br>
1:使用[https://cloud.tencent.com/developer/article/1388556]
</br></br>
`1万属性，100亿数据，每秒10万吞吐，架构如何设计？`
</br>
1:使用[
https://cloud.tencent.com/developer/article/1420239]
</br></br>
`Windows10 + WSL`
</br>
1:使用[https://www.jianshu.com/p/8404e34feefe]
</br>
2:使用[https://blog.csdn.net/qq_36620040/article/details/90301526]
</br></br>
`基于Redis实现排行榜周期榜与最近N期榜`
</br>
1:使用[https://www.jianshu.com/p/8404e34feefe]
</br></br>
`爬去天天基金网接口`
</br>
1:使用[https://www.jianshu.com/p/d79d3cd62560]  或者 [https://www.jianshu.com/p/f0607374b8ee?from=timeline&isappinstalled=0]
</br>
2:使用[https://blog.csdn.net/chaoren499/article/details/90232332]
</br></br>
`JS纯前端实现audio音频剪裁剪切复制播放与上传`
</br>
1:使用[https://www.zhangxinxu.com/wordpress/2020/07/js-audio-clip-copy-upload/]
</br></br>
`TradingView + WebSocket 实时推送 K 线脱坑指南`
</br>
1:使用[https://juejin.im/post/5bb1f22151882557305835f8]
</br></br>
`Redis中BitMap是如何储存的，以及PHP如何处理`
</br>
1:使用[https://www.cnblogs.com/sunlong88/articles/9251645.html]
</br></br>
`AC自动机-去除敏感字符`
</br>
1:使用[https://www.jianshu.com/p/9f4271fcbc76]
</br></br>
`GO 条件锁sync.Cond`
</br>
1:使用[https://www.jianshu.com/p/2dce206b6d79]
</br></br>

`PACING算法 广告预算 浅谈广告系统预算控制`
</br>
1:使用[https://www.lizenghai.com/archives/42609.html]
</br></br>

`singleflight  golang缓存防击穿庫`
</br>
1:使用[https://github.com/golang/groupcache/singleflight/]
</br></br>


## 数据分析思维
### 七技巧
### 三核心
### 案例分析
根据多个维度来加工出一个综合评价系数</br>
指标有:点评、人均、口味、环境、服务</br>
步骤:</br>
1.每个指标下的值有可能存在空值,应该先将带空值的数据清洗掉或者给个平均数填充</br>
2.将数值差异较大的数据也清洗掉：如点评指标中存在一个2w多的数据而其他都是8k左右的，就把2w的数据清掉</br>
3.分析指标：如口味、环境、服务代表餐厅质量，点评代表热度，前面的指标是正比关系，而人均指标属于降权(多数人希望物美价廉)</br>
4.开始把指标构成公式：</br>
   4.1此处将口味、环境、服务的值累加，如果更看重口味就给口味的值加个权重系数-->公式：(1.2*口味+环境+服务)/3</br>
   4.2因为点评指标的值存在很大的差异化(最大值和最小值差异很多)，因此使用指数法来进行指标值的收敛-->log(点评,2)，如果底数为2计算出的结果还是存在挺大的差异，可以更改底数为10或其他，尽量让计算出来的结果差异上缩小(如缩到0.n的差异)</br>
   4.3人均指标也一样，可以使用指数法来进行收敛：此处的底数代表用户对价格的敏感度，底数越大越不敏感，像以底数为10的话人均100和人均1000计算出的结果相差不大的，应该尽量避免，此处使用4作为底数log(人均,4)</br>
   4.4将上面3个步骤计算的数值进行计算:综合评价系数=[(1.2*口味+环境+服务)/3]+log(点评,2)-log(人均,4)</br>
   4.5上面计算出来的结果未必靠谱，因此接下来使用统计学的MinMax Scaler方法(归一化):</br>
        4.5.1:先计算出综合评价系数中的最大值MAX(综合评价系数)和最小值MIN(综合评价系数)，将最大值和最小值进行相减MAX(综合评价系数)-MIN(综合评价系数)</br>
        4.5.2:将数据收敛至0到1之间:(int-min)/(max-min)=(综合评价系数-MIN(综合评价系数))/MAX(综合评价系数)-MIN(综合评价系数)</br>
以上公式的人均指标计算使用的底数为4，权重倾向大，可根据实际情况更改权重底数</br>
</br></br>



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
2: 网盘搜索 [https://dalipan.com/search?keyword=] </br> [https://www.iwanw.com/] </br>
3: 短视频去水印 [https://imyshare.com/parsevideo/] </br> [https://dy.kukutool.com/] </br> [http://douyin.haiya360.com/] </br> [https://douyin.video996.com/] </br> [https://www.videoparse.cn/] </br> 
4: 电子书搜索 [https://www.jiumodiary.com/] </br>  [http://www.lingocn.com/category/ebook] </br> 
5: 图片压缩 [https://works.yangerxiao.com/icfe/] </br>
6: 百度文库在线解析 [http://www.html22.com/doc/] </br>  [https://www.docin365.com/] </br>
7: 视频解析 [http://weibodang.cn/] </br> [https://tool.lu/videoparser/] </br>  
8: 多功能在线工具(文件格式转换等) [https://zh.pickfrom.net/] </br>
9: 下载youtube字幕文件 [https://downsub.com/] </br>

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
6:合并视频1【将要合并的视频名字写在同级的filelist.txt文件中，格式为 file '1.mp4'】：`ffmpeg -f concat -i filelist.txt -c copy output.mp4`
</br></br>
7:合并视频2【如合并的视频出现画音轨不统一，则使用该方法:先将文件转成ts格式，再把ts合并】：`ffmpeg -i output1.mp4 -vcodec copy -acodec copy -vbsf h264_mp4toannexb output1.ts`
</br>
`ffmpeg -f concat -i filelist.txt -acodec copy -vcodec copy -absf aac_adtstoasc output.mp4`
</br></br>
8:截取制定位置的视频【将视频2.mp4，从5秒开始剪辑时长为4分41秒的视频output.mp4】：`ffmpeg -ss 00:00:05 -t 00:04:41 -i 2.mp4 -vcodec copy -acodec copy output.mp4`
</br></br>
9:裁剪视频宽高画面【将视频1.mp4的画面高度减去原来高度的1.2成视频output.mp4】：`ffmpeg -i 1.mp4 -strict -2 -vf crop=iw:ih/1.2:0:ih out1.mp4`
</br></br>
10:将视频的音频抽出:`ffmpeg -i we.mp4 -vn r.mp3`
</br></br>
11:将视频宽高改变:`ffmpeg -i we.mp4 -s 720*1280 testp.mp4`
</br></br>
12:将视频添加文字水印:`ffmpeg -i we.mp4 -vf "drawtext=fontcolor=black:fontfile=simsun.ttc:text='喜欢就关注一下':x=20:y=20:fontsize=40:shadowy=2" outpu.mp4`
</br></br>
13:预览消除水印位置:`ffplay -i i.mp4 -vf delogo=x=1:y=1:w=168:h=30:show=1`
</br></br>
14将视频水印模糊:`ffmpeg -i we.mp4 -vf "delogo=x=490:y=0:w=136:h=64" -c:a copy newp.mp4`
</br></br>
15视频截取指定大小的图片:`ffmpeg -i i.mp4 -y -f image2 -ss 00:00:26 -t 0.001 -s 1366x768 a.jpg`
</br></br>
16给视频加上字幕:`ffmpeg -i i.mp4 -vf subtitles=word.srt -y output.mp4`
</br></br>


## SQL优化
1:【分页查询优化】在分页查询大量数据时候，采用主键索引来定位快,一般的查询采用limit是会逐一统计数据条数。如SELECT * FROM table WHERE id>500000 AND id<=500000+100 的查询速度会比一般分页查询SELECT * FROM table LIMIT 500000,100。
</br></br>
2：https://developer.aliyun.com/article/727673?utm_content=g_1000089764
</br></br>
3:【查询同时包含多个指定标签的文章】查询 tag_id 同时包含 1、2、3 的 article_id：</br>
```sql
SELECT article_id
FROM articles_tags
WHERE tag_id in (1, 2, 3)
GROUP BY article_id
HAVING COUNT(*) = 3
```
</br></br>
4:【一次性更新多值】更新id为13,12,11其所对应的age：</br>
```sql
UPDATE table SET age = CASE id 
   WHEN 13 THEN 45 
   WHEN 12 THEN 34 
   WHEN 11 THEN 87 
END WHERE id IN (13,12,11)
```
</br></br>
5:【一个月按7天作为周维度查询】查询2020年8月(15962112000-1598889599)每周支付人数和支付金额：</br>
```sql
SELECT
    FLOOR((day(FROM_UNIXTIME(stat_time,'%Y-%m-%d'))-1)/7)+1 weekapn,sum(money) money,sum(pay_num) paynum
FROM t_data where stat_time BETWEEN 1596211200 and 1598889599 group by weekapn order by weekapn
```
</br></br>
6:【查询组内Top N:】查询求出各部门薪水排名前2的员工，如果有多个员工薪水相同，需要全部显示：</br>
```sql
select * from salaries s1 where (select count(distinct(s2.salary)) from salaries s2 where s1.dept_no = s2.dept_no and s1.salary < s2.salary) < 2;
+----+--------+---------+--------+
| id | emp_no | dept_no | salary |
+----+--------+---------+--------+
|  1 |      1 |       1 |  10000 |
|  2 |      2 |       1 |  10000 |
|  3 |      3 |       1 |   8000 |
|  4 |      4 |       1 |   8000 |
|  6 |      6 |       2 |  10000 |
|  7 |      7 |       2 |   9000 |
|  9 |      9 |       3 |  10000 |
| 10 |     10 |       3 |   9000 |
| 11 |     11 |       3 |   9000 |
+----+--------+---------+--------+
9 rows in set (0.00 sec)
```
```
这语句要怎么理解呢？首先，如何定义薪水排名第1，换个说法就是薪水比我高的一个都没有，有0个。如何定义薪水排名第2，就是在薪水在我之上的，薪水去重后值只是1个，依此类推。上边语句把最末的”< 2″换成”in(0, 1)”估计就更好理解了。Top N就是要把排名第1，第2，…第N的取回来，in (0, 1, … , N – 1)，换个简洁的写法”< N”。因此，如果把问题改为排名第N的，很显然，结果就是”in(N – 1)”。
```
</br></br>
7:【MySQL或MariaDB，不要用“utf8”编码，改用“utf8mb4”】避免特殊字符缺失或者查询特殊字符报错导致链接超时问题：[https://mathiasbynens.be/notes/mysql-utf8mb4#utf8-to-utf8mb4]
</br></br>


## Golang环境安装下载


## Golang环境安装下载
1:下载源码包【此处我下载的是最新的版本】：`sudo curl -OL https://studygolang.com/dl/golang/go1.15.src.tar.gz`
</br></br>

## Golang调用系统默认浏览器打开指定链接
来源：[https://www.jianshu.com/p/29adf056e72b]
</br></br>
1.windows
```go
cmd := exec.Command(`cmd`, `/c`, `start`, `http://baidu.com`)
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
err :=cmd.Start()
if err!=nil{
   fmt.Println(err)
}
```
## Golang设置程序图标
来源：[https://studygolang.com/articles/29850?fr=sidebar] </br> [https://blog.csdn.net/u014633966/article/details/82984037] </br>


## go.mod的使用，告别GOPATH
来源：[https://www.jianshu.com/p/bbed916d16ea]
1：在当前目录下新建一个go文件，里面直接调用从未下载过得github包，保存退出
</br></br>
2：在当前目录使用命令`go mod init wserver`（go mod init 后面需要跟一个名字，我这里叫wserver）,看到提示 “go: creating new go.mod: module wserver”  说明 go mod 初始化成功了，会在当前目录下生成一个 go.mod 文件。
</br></br>
3:然后直接run一下新建的go文件：`go run main.go`，即可发现自动下载里面的包，并且正常运行了程序。同时项目目录下多出了一个文件 go.sum（go.sum 是记录所依赖的项目的版本的锁定）
</br></br>
4:如果出现【connect: connection refused】代表自动下载包失败，需要设置一下go的环境变量配置(被墙了)：使用命令`go env -w GOSUMDB=off`关闭到墙的下载设置，然后再使用命令`go env -w GOPROXY=https://goproxy.cn`设置国内的包下载代理，再重新`go run main.go`一下，应该可以正常运行了

## 熔断判断公式
dropRatio=max(0,[(request-protetion)-K*accepts]/request+1)

## 计算请求平均耗时百分比
```go
func ordinal(length int, percentile float64) int64 {
   if percentile == 0 && length > 0 {
	return 1
   }
   return int64(math.Ceil((percentile / float64(100)) * float64(length)))
}

func Percentile(p float64)  {
	//一段时间内的请求耗时
	sortedDurations := []time.Duration{3000000,2000000,30000000,4000000,5000000}
	//length := len(sortedDurations)
	pos := ordinal(len(sortedDurations), p) - 1
	//单位转成ms
	return uint32(sortedDurations[pos].Nanoseconds() / 1000000)
}

func main(){
   //计算99.5%的请求耗时
   fmt.Println(Percentile(99.5))
}
```
## GO的常用代码段
1：头条SDK的key加密类型和发送post请求
</br>
```golang
const (
   AppKey = ""
   APP_API_KEY = ""
)
func getDataSignWithKey(paramsMap map[string]interface{},secretKey string) string{
	appkey := AppKey
	//先获取所有键名
	var keys []string
	for k := range paramsMap{
	    keys = append(keys, k)
	}
	//将key按照字段序排序
	sort.Strings(keys)
	// 生成待签名数据
	buf := new(bytes.Buffer)
	for _, k := range keys {
	    buf.WriteString(fmt.Sprintf("%s=%v", k, paramsMap[k]))
	}
	buf.WriteString(fmt.Sprintf("%v%v",appkey,secretKey))
	// 生成MD5签名值
	return fmt.Sprintf("%x", md5.Sum(buf.Bytes()))
}

func main(){
	apiUrl := "http://https://ad.oceanengine.com//api/login.php"
	fieldMap := map[string]interface{}{ // 未签名数据
		"req_time":   time.Now().Unix() ,
		"account_name":     fmt.Sprintf("testAccount%v",time.Now().Unix()),
	}
	
	postMaps := make(map[string]interface{})
	data,_:=json.Marshal(fieldMap)
	//urldecode解码
	//enEscapeUrl, _ := url.QueryUnescape(escapeUrl)
	//urlencode编码
	postMaps["data"] = url.QueryEscape(string(data))
	postMaps["signcode"] = getDataSignWithKey(fieldMap,APP_API_KEY)
	
	
	//组装发送post表单请求
	buf := new(bytes.Buffer)
	for k, v := range postMaps {
	    buf.WriteString(fmt.Sprintf("%s=%v&", k, v))
	}

	req, _ := http.NewRequest(http.MethodPost, apiUrl, bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//如果是发送json的话
	//req.Header.Set("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
	    // handle error
	    fmt.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```
</br></br>
2：头条接口的发送GET请求
</br>
```golang
func main(){
	apiUrl := "https://ad.oceanengine.com/track/activate/"
	payloadMap := map[string]interface{}{ // 未签名数据
		"callback":   "EJiw267wvfQCGKf2g74ZIPD89-vIATAMOAFCIjIwMTkxMTI3MTQxMTEzMDEwMDI2MDc3MjE1MTUwNTczNTBIAQ%3D%3D",
		"conv_time":  "1574835097",
		"event_type": "3",
		"imei":       "0c2bd03c39f19845bf54ea0abafae70e",
	}
	//发送GET请求
	req, _ := http.NewRequest(http.MethodGet, apiUrl, nil)
	q := req.URL.Query()
	for k,v := range payloadMap{
	    q.Add(k,v.(string))
	}
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {// handle error}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```
</br></br>
3：Go类型递归构建字典树
</br>
```golang
package main

import (
	"fmt"
	"unsafe"
)

//字典树
type (
	Char  = rune
	Trie map[Char]Trie
)

//添加字符串
func (t Trie) AddString(s string)  {
	for _,ch := range []Char(s){
		if t[ch]!=nil {
			t=t[ch]
		}else {
			prev := t
			t = make(Trie)
			prev[ch]=t
		}
	}
}

//判断字符串
func (t Trie) ExistPrefix(prefix string)bool{
	for _,ch:=range []Char(prefix){
		if t[ch] == nil{
			return false
		}
		t = t[ch]
	}
	return true
}

//查询关联推荐
func (t Trie) Suggest(prefix string)[]string{
	 prefixBytes := []Char(prefix)
	 for _,ch := range prefixBytes{
	 	t=t[ch]
	 }
	 return dfs(t,prefixBytes,nil)
}

//辅助函数
func dfs(t Trie,prefix []Char,suggestions []string)[]string{
	for ch,newT:=range t{
		prefix:=append(prefix,ch)
		if len(newT)==0 {
			suggestions = append(suggestions,string(prefix))
		}else {
			suggestions = dfs(newT,prefix,suggestions)
		}
	}
	return suggestions
}


func main() {
	str :="我在"
	str2:="我在这"
	str3:="我在哪"

	t := make(Trie)
	t.AddString(str)
	t.AddString(str2)
	t.AddString(str3)

	fmt.Println(
		t.ExistPrefix("我在这"),//判断该字符串是否存在
		t.Suggest("我在"),//print [我在这，我在哪]
		unsafe.Sizeof(t))//8个字节的占用大小
}
```
</br></br>


## 二进制求集合
1：【交集】`∩交集:$ c = $value1 & $value2`
</br></br>
2：【差集】`差集:$ c= ~( $value1 | $value2 )`
</br></br>
3：【并集】`∪并集:$c = ( $value1 | $value2 )`
</br></br>

## 利用差分数组和前缀和来统计每个位置的出现次数
```golang
func demo(){
	//统计每个位置的出现次数
	cnt := make([]int,5)
	fmt.Println(cnt)
	//{i,j}表示数组的第i位到第j位的都出现过一次
	request:=[][]int{{1,3},{0,1}}
	for _,v:=range request{
		//差分数组
		//开始位
		cnt[v[0]]+=1
		//结束位
		cnt[v[1]+1]-=1
	}
	fmt.Println(cnt)

	//前缀和
	for i:=1;i<n;i++{
		cnt[i]+=cnt[i-1]
	}
	fmt.Println(cnt)
}
```


## MYSQLDump-tips
1:【导出不影响：mysqldump --single-transaction】--single-transaction参数的作用，设置事务的隔离级别为可重复读，即REPEATABLE READ，这样能保证在一个事务中所有相同的查询读取到同样的数据，也就大概保证了在dump期间，如果其他innodb引擎的线程修改了表的数据并提交，对该dump线程的数据并无影响，在这期间不会锁表。

## xmSelect常用法
```js
   <div class="layui-inline">
         <div class="layui-input-inline" id="selectSonChannel"></div>
    </div>
   //直接渲染
   xmSelect.render({
            el: '#selectSonChannel',
            radio: true,
            empty: '呀, 没有数据呢',//没数据的提示
            name:'channel_id',//发送到后端的字段
            tips:'请选择二级渠道',
            clickClose:true,
            filterable: true,
            prop: {value: 'value', name: 'name'},//此处的属性跟下面的data属性对应
            data:[{'name':'test','value':123}]
        });
   //延迟渲染
   let dom = xmSelect.render({
            el: '#selectSonChannel',
            radio: true,
            empty: '呀, 没有数据呢',//没数据的提示
            name:'channel_id',//发送到后端的字段
            tips:'请选择二级渠道',
            clickClose:true,
            filterable: true,
            prop: {value: 'value', name: 'name'},//此处的属性跟下面的data属性对应
            data:[]
        });
	
   res = getPostFromServerData()
   
   dom.update({
            data:res.data.map(function(value,index,array){
                 return {'name':value.sub_channel_name,'value':value.channel_id}
            })
    })
    
    //选择后监听
    xmSelect.render({
            el: '#selectSonChannel',
            radio: true,
            empty: '呀, 没有数据呢',//没数据的提示
            name:'channel_id',//发送到后端的字段
            tips:'请选择二级渠道',
            clickClose:true,
            filterable: true,
            prop: {value: 'value', name: 'name'},//此处的属性跟下面的data属性对应
            data:[{'name':'test','value':123}],
	    on:function(data){
	    	 //arr:  当前多选已选中的数据
                 var arr = data.arr;
		 if(arr[0]){//TODO::}
	    },
        });
	
   //从服务器获取数据
   xmSelect.render({
            el: '#selectSonChannel',
            radio: true,
            paging: true,
            pageRemote: true,
            clickClose:true,
            filterable: true,
            prop: {value: 'value', name: 'name'},
            remoteMethod: function(val, cb, show, pageIndex){
	    	 res = getPostFromServerData()
		 cb(res.data.map(function(value,index,array){
                       return {'name':value.name,'val':value.id}
                  }), res.total_page);
	    }
        });
   
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
array_multisort(array_column($b, 'id'), SORT_DESC, SORT_NUMERIC,$b);
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


//获取文件格式
echo pathinfo("record.log", PATHINFO_EXTENSION); //log

//获取图像信息
list($width, $height, $type, $attr) = getimagesize("runoob-logo.png");
echo "宽度为：" . $width;
echo "高度为：" . $height;
echo "类型为：" . $attr;

//汉子拆分成数组
var_dump(preg_split('//u',$str,null,PREG_SPLIT_NO_EMPTY))


//使用关联数组来取代if-else/switch
$res = ['over_45'=>0,'over_30'=>0,'over_15'=>0];
$data=[
    46,39,29,17,
];

function judge($condition,&$da){
    $map = [
        ($condition > 15) =>function()use(&$da){
            return $da['loss_15']++;
        },
        ($condition > 30) =>function()use(&$da){
            return $da['loss_30']++;
        },
        ($condition > 45) =>function()use(&$da){
            return $da['loss_45']++;
        },
    ];
    //因为在声明的时候就自动判断并且保留下条件为真的状态
    $map[1]();
}

foreach ($data as $v){
    test($v,$res);
}

var_dump($res);//['over_45'=>1,'over_30'=>1,'over_15'=>2];


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




//2-图片base64编码在服务端保存
public function base64ImageContent($base64){
        if (preg_match('/^(data:\s*image\/(\w+);base64,)/', $base64, $result)){
            $type = $result[2] == 'jpeg' ? 'jpg' : $result[2];
            $path = '/'.date('Ym').'/';
            $img_name = date('dHis') . uniqid() .'.'. $type;
            if (!is_dir(dirname($this->nfs.'/image/'.$path))) {
                mkdir(dirname($this->nfs.'/image/'.$path), 0777, true);
            }
            if (file_put_contents($this->nfs.'/image/'.$path . $img_name, base64_decode(str_replace($result[1], '', $base64)))){
                return $path . $img_name;
            }else{
                return false;
            }
        }else{
            return false;
        }
    }
   
   
//3-使用Client上传文件
$client = new Client();
        $options = [
            'timeout' => 120,
            'verify' => false,
            'headers' => [
                'Access-Token' =>'',
            ],
            'query' => [],
            'multipart' => [
                [
                    'name' => 'file',
                    'contents' => fopen($params['file'], 'r'),
                ],
                [
                    'name' => 'signature',
                    'contents' => md5_file($params['file']),
                ],
            ],
        ];
        $res = $client->request('POST', 'url', $options);
        return json_decode($res->getBody(), true);
	
	
//4-递归思想实现中间件
class Middlewares {
    public $middlewares=[];
    public $index=-1;
    public $request;

    public function __construct($mid)
    {
        $this->middlewares=$mid;
    }

    public function next()
    {
        $this->index++;
        $s=count($this->middlewares);
        for (;$this->index<$s;$this->index++){
            $this->middlewares[$this->index]($this);
        }
    }
}

   $c=new Middlewares([
       function(Middlewares $class){
           LogLib::write([
               1
           ], 'time');
           $class->next();
           LogLib::write([
               4
           ], 'time');
       },
       function(Middlewares $class){
           LogLib::write([
               2
           ], 'time');
           $class->next();
           LogLib::write([
               3
           ], 'time');
       },
       function()use($application){
           LogLib::write([
               'hello'
           ], 'time');
           echo $application->handle()->getContent();
       }
   ]);

   $c->next();
   //1，2，hello，3，4



//5-pipeline管道处理流水逻辑，可以参考:https://learnku.com/articles/38189
//逻辑处理接口
interface  TestUnit
{
    public function handle($passable, callable $next = null);
}
//逻辑1
class  Unit1 implements TestUnit
{
    public function handle($passable, callable $next = null)
    {
        echo __CLASS__ . '->' . __METHOD__ . " called\n";
        $next($passable);
    }
}
//逻辑2
class  Unit2 implements TestUnit
{
    public function handle($passable, callable $next = null)
    {
        echo __CLASS__ . '->' . __METHOD__ . " called\n";
        $next($passable);
    }
}
//逻辑3
class  Unit3 implements TestUnit
{
    public function handle($passable, callable $next = null)
    {
        echo __CLASS__ . '->' . __METHOD__ . " called\n";
        $next($passable);
    }
}
//逻辑4
class  InitialValue implements TestUnit
{
    public function handle($passable, callable $next = null)
    {
        echo __CLASS__ . '->' . __METHOD__ . " called\n";
        //$next($passable);
    }
}
//使用类的逻辑处理顺序
$processObj = [new Unit1(), new Unit2(), new Unit3()];
//使用闭包的逻辑处理循序
$processCallback = [
    function($passable, callable $next = null){echo 'I am One:'.$passable.PHP_EOL;$passable++;$next($passable);},
    function($passable, callable $next = null){echo 'I am Two:'.$passable.PHP_EOL;$passable++;$next($passable);},
    function($passable, callable $next = null){echo 'I am Three:'.$passable.PHP_EOL;$passable++;$next($passable);},
];
//此处是管道处理的主体函数
function Pipeline($stack, $pipe)
{
    //此处的$passable是由外部传入的数据
    return function ($passable) use ($stack, $pipe) {
    	//$pipe是逻辑的实体，$stack紧跟着的下一个逻辑的处理体

        if (is_callable($pipe)) {
            $pipe($passable, $stack);
        } elseif (is_object($pipe)) {
            $method = "handle";
            if (!method_exists($pipe, $method)) {
                throw new InvalidArgumentException('object that own handle method');
            } else {
                $pipe->$method($passable, $stack);
            }
        } else {
            throw new InvalidArgumentException('$pipe must be callback or object');
        }
    };
}
$pipeline = array_reduce(array_reverse($processCallback), "Pipeline", function ($passable) {
    (new InitialValue())->handle($passable);
});
//设置要传入处理的数据
$data = 90;
$pipeline($data);
//结果是依次执行
//I am One:90
//I am Two:91
//I am Three:92
//InitialValue->InitialValue::handle called



//6-使用闭包替代递归函数
$fib = function ($n) use (&$fib) {
    if ($n == 0 || $n == 1) {
        return 1;
    }

    return $fib($n - 1) + $fib($n - 2);
};
$fib(3)//非波拉契数列

//7-使用生成器yield优化内存性能(只配合foreach来遍历)
function createRange($number){
    for($i=0;$i<$number;$i++){
        yield time();
    }
}
//无论开始传入的 $number 有多大，由于并不会立即生成所有结果集，所以内存始终是一条循环的值。
foreach(createRange(10000000000) as $value){
    echo $value.'<br />';
}

//通常方法，如果是百万级别的访问量，会占用极大内存  ,改用yield就不会全加载，只会逐个使用
function rand_weight($numbers)
{
    $w = function($numbersArr){
        $total = 0;
        foreach ($numbersArr as $number => $weight) {
            $total += $weight;
            yield $number=>$total;
            //$distribution[$number] = $total;
        }
    };
    $total = array_sum($numbers);
    $rand = mt_rand(0, $total-1);
    var_dump($rand);

    foreach ($w($numbers) as $num => $weight) {
        if ($rand < $weight) return $num;
    }
}  
$numbers = array('nike' => 200, 'jordan' => 500, 'adiads' => 800);  
var_dump(rand_weight($numbers));

//PHP开发很多时候都要读取大文件，比如csv文件、text文件，或者一些日志文件。这些文件如果很大，比如5个G。这时，直接一次性把所有的内容读取到内存中计算不太现实。这里生成器就可以派上用场啦。
function readTxt()
{
    # code...
    $handle = fopen("./test.txt", 'rb');

    while (feof($handle)===false) {
        # code...
        yield fgets($handle);
    }

    fclose($handle);
}

foreach (readTxt() as $key => $value) {
    # code...
    echo $value.'<br />';
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


## js小功能

```js
//1：复制到剪贴板
function copy() {
        const input = document.createElement('input');
        const value = 'value';
        input.setAttribute('readonly', 'readonly');
        input.setAttribute('value', value);
        document.body.appendChild(input);
        input.setSelectionRange(0, value.length);
        input.select();
        if (document.execCommand('copy')) {
            document.execCommand('copy');
            alert('礼包码已复制到粘贴板');
        }
        document.body.removeChild(input);
    }
    
    
//2：获取url上的参数
var qs = (function(a) {
            if (a == "") return {};
            var b = {};
            for (var i = 0; i < a.length; ++i)
            {
                var p=a[i].split('=', 2);
                if (p.length == 1)
                    b[p[0]] = "";
                else
                    b[p[0]] = decodeURIComponent(p[1].replace(/\+/g, " "));
            }
            return b;
        })(window.location.search.substr(1).split('&'));
console.log(qs['user_name']);	


//3：本地上传的图片预览显示
<div>
    <input type="file" id="p" onchange="show(this)">
    <img src="" alt="" id="pic">
</div>
<script>
    function show(self) {
        var reader= new FileReader();
        reader.readAsDataURL(self.files[0]);
        reader.onload=function(e){
            document.getElementById("pic").src = e.target.result;
        }
    }
</script>

//4：视频video截图
//html标签如下：
 <div>
    <video id="video" autoplay="true"  width="500" height="400" src="http://vjs.zencdn.net/v/oceans.mp4"></video>
    <canvas id="canvas" width="500" height="400" ></canvas>
    <button id="btn">截取</button>
 </div>
 <script>
 var canvas = $("#canvas"),
  ctx = canvas[0].getContext("2d"),
  width = $("#video").width(),
  height = $("#video").height(),
  drawImg = function(){
        ctx.drawImage($("#video")[0], 0, 0, width, height);
	//输出画布图像的base64格式(注意！如果video的链接跟当前域名非同源的话是会报错的)
        console.log(document.getElementById("canvas").toDataURL());
    }
  $("#btn").on("click", function(){
        drawImg();
   });
 </script>


//5：上拉刷新
<script>
$(window).scroll(function(){
   if( ($(window).height() + $(window).scrollTop()) > $(document).height() - 20 ){
   	//TODO::ajax load data
   }
});
</script>


//6：base64编码解码
  //编码
btoa('hello')
  //解码
atob('aGVsbG8=')


//7：笛卡尔积实现数组元素全排列
function cartesianProduct (arr) {
  if (arr.length == 1) {
    return arr;
  } else {
    return arr.reduce(function (a, b) {
      return a.map(function (x) {
        return b.map(function (y) {
          return x.concat(y);
        })
      }).reduce(function (a, b) { return a.concat(b) }, [])
    }, [[]])
  }
}

console.log(cartesianProduct([[1,2,3],['a','b']]));
//0: (2) [1, "a"]
//1: (2) [1, "b"]
//2: (2) [2, "a"]
//3: (2) [2, "b"]
//4: (2) [3, "a"]
//5: (2) [3, "b"]

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
## window-API函数
具体参考文章[https://www.cnblogs.com/dfsxh/archive/2010/11/29/1890984.html#2007524]

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

## itemCF-demo
详细相关可以参考文章[https://blog.csdn.net/Smart3S/article/details/86764297] 
```php
// 从文件中获取数据流
function get_user_click($rating_file){
    //存储用户选择过的item容器
    $user_click = [];
    //存储用户选择过的item的时间戳容器
    $user_click_time = [];
    $file = fopen($rating_file, "r");
    while(! feof($file))
    {
        $line =  fgets($file)."\n";//fgets()函数从文件指针中读取一行
        $a = explode(",",$line);
        list($user_id,$item_id,$rating,$timestamp) = $a;
        //评分小于3.0则为不喜欢，数据跳过
        if ($rating<3.0) {
            continue;
        }
        if(!isset($user_click[$user_id])){$user_click[$user_id] = [];}
        //存储用户选择过的item
        $user_click[$user_id][] = $item_id;
        //存储选择该item时的时间戳
        if(!isset($user_click_time[$user_id.'_'.$item_id])){$user_click_time[$user_id.'_'.$item_id] = $timestamp;}
    }
    fclose($file);
    return [$user_click,$user_click_time];
}
// 计算惩罚时间差越大的用户的权重
function base_contribute_score($click_time_i,$click_time_j){
    $click_time_i = floatval($click_time_i);
    $click_time_j = floatval($click_time_j);
    $delata_time =abs($click_time_i-$click_time_j);
	//将时间戳单位换算为天
	$total_sec=60*60*24;
	$delata_time=$delata_time/$total_sec;
	return 1/(1+$delata_time);
}
// 计算相似度
function cal_item_sim($user_click,$user_click_time){
    //存储item被用户的选择次数
    $item_user_click_count = [];
    //存储相似度数据
    $co_appear = [];
    //存储item的相似度得分
    $item_sim_score = [];
    //循环选择序列数据，user是每个用户的id，itemlist是每个用户的选择序列
    foreach ($user_click as $user=>$itemList){
        foreach ($itemList as $index_i=>$itemId_i){
            if (!isset( $item_user_click_count[$itemId_i])){$item_user_click_count[$itemId_i] = 0;}
            $item_user_click_count[$itemId_i]+=1;
            //计算每个item的id和其他item的id共同出现在一个用户的选择序列中的数值
            $ts = $itemList;
            //删除第一个元素
            array_shift($ts);
            foreach ($ts as $index_j=>$itemId_j){
                //获取点击的时间戳
                //增加时间衰减因子对推荐的考虑
                $click_time_i=0;
                $click_time_j=0;
                if (isset( $user_click_time[$user.'_'.$itemId_i])){$click_time_i = $user_click_time[$user.'_'.$itemId_i];}
                if (isset( $user_click_time[$user.'_'.$itemId_j])){$click_time_j = $user_click_time[$user.'_'.$itemId_j];}
                //计算所有item的id中，两两id的共同出现次数
                if (!isset( $co_appear[$itemId_i])){$co_appear[$itemId_i] = [];}
                if (!isset( $co_appear[$itemId_i][$itemId_j])){$co_appear[$itemId_i][$itemId_j] = 0;}
                $co_appear[$itemId_i][$itemId_j]+=base_contribute_score($click_time_i,$click_time_j);
                if (!isset( $co_appear[$itemId_j])){$co_appear[$itemId_j] = [];}
                if (!isset( $co_appear[$itemId_j][$itemId_i])){$co_appear[$itemId_j][$itemId_i] = 0;}
                $co_appear[$itemId_j][$itemId_i]+=base_contribute_score($click_time_i,$click_time_j);
            }
        }
    }
    //相似度计算
    foreach ($co_appear as $itemId_i=>$relate_item){
        foreach ($relate_item as $itemId_j=>$co_time){
            //相似度计算
            //原数学公式为：(选择i的数量集数与选择j的数量集数)的交集/[(选择i的数量集数与选择j的数量集数)的并集]开根号
            $sim_score = $co_time / sqrt(floatval($item_user_click_count[$itemId_i]*$item_user_click_count[$itemId_j]));
            if (!isset( $item_sim_score[$itemId_i])){$item_sim_score[$itemId_i] = [];}
            $item_sim_score[$itemId_i][$itemId_j] = $sim_score;
        }
    }
    return $item_sim_score;
}
// 对相似度进行降序排序
function sore_item_sim($sim_info){
    foreach( $sim_info as $itemid=>$value){
        arsort($value);
        $sim_info[$itemid] = $value;
    }
    return $sim_info;
}
// 计算每个用户的推荐（与喜好item相似度较高的）物品
function cal_recom_result($sim_info,$user_click){
    $recom_info = [];
    //选取用户的前三个选择item作为item样本，找寻它们的相似item进行推荐
    $recent_click_num = 3;
    //取item样本中电影相似度前topk的电影
    $topk = 5;
    foreach ($user_click as $user=>$click_list){
        if(isset($recom_info[$user])){$recom_info[$user] = [];}
        $ts = array_slice($click_list,0,$recent_click_num);
        foreach ($ts as $item_id){
            if(!isset($sim_info[$item_id])){continue;}
            //找寻与item相似度前topk的item
            $tp = array_slice($sim_info[$item_id],0,$topk);
            foreach ($tp as $itemsimid=>$itemsimscore){
                $recom_info[$user][$itemsimid] = $itemsimscore;
            }
        }
    }
    return $recom_info;
}
function main_flow(){
    list($user_click,$user_click_time) = get_user_click('p.txt');
    $sim_info = sore_item_sim(cal_item_sim($user_click,$user_click_time));
    $recom_info = cal_recom_result($sim_info,$user_click);
    var_dump($recom_info["1"]);    
}
main_flow();
```

## Trie树关键词搜索
```php
class Trie
{
    /**
     * node struct
     *
     * old version
     * node = array(
     * val->word
     * next->array(node)/null
     * depth->int
     * )
     *
     * new version
     * node = [val,next]
     */
    private $root = [];
    private $matched = [];
    public function append($keyword)
    {
        $words = preg_split('/(?<!^)(?!$)/u', $keyword);
        array_push($words, '`');
        $this->insert($this->root, $words);
    }
    public function match($str)
    {
        $this->matched = [];
        $words = preg_split('/(?<!^)(?!$)/u', $str);
        while (count($words) > 0) {
            $matched = [];
            $res = $this->query($this->root, $words, $matched);
            if ($res) {
                $this->matched[] = implode('', $matched);
            }
            array_shift($words);
        }
        return $this->matched;
    }
    private function insert(&$node, $words)
    {
        if (empty($words)) return;
        $word = array_shift($words);
        if (isset($node[$word])) {
            $this->insert($node[$word], $words);
        } else {
            $word == '`' ? $node[$word] = 1 : $node[$word] = [];
            $this->insert($node[$word], $words);
        }
    }
    private function query($node, $words, &$matched)
    {
        $word = array_shift($words);
        if (isset($node[$word])) {
            array_push($matched, $word);
            if (isset($node[$word]['`'])) {
                return true;
            }
            return $this->query($node[$word], $words, $matched);
        } else {
            $matched = array();
            return false;
        }
    }
}
$node = new Trie();
//存入关键词(屏蔽词)
$node->append("科学");
$node->append("科学家");
$node->append('我们');
//匹配字符串中存在的关键词(屏蔽词)
$res = $node->match("我们都是科学家,你们说的对不对？我们觉得对");
var_dump($res);
```

## 图的知识-Graph
```php
//稠密图-邻接矩阵
////0 1 2
//0 f t f    0<-->1(无向图) 0->1(有向图)
//1 t f f
//2 f f f
class DenseGraph{
    private $n,$m,$directed,$g=[];
    public function __construct($n,$directed)
    {
        $this->n=$n;//定点个数
        $this->m=0;//初始边都为0
        $this->directed=$directed;//是否有向(是a->b还是a<->b)
        for ($i=0;$i<$n;$i++){
            array_push($this->g,array_fill(0,$n,false));
        }
    }
    public function V(){return $this->n;}
    public function E(){return $this->m;}
    private function hasEdge($v,$w){
        if ($v<0 || $v>=$this->n){return false;}
        if ($w<0 || $w>=$this->n){return false;}
        return $this->g[$v][$w];
    }
    public function addEdge($v,$w)
    {
        if ($v<0 || $v>=$this->n){return false;}
        if ($w<0 || $w>=$this->n){return false;}
        if ($this->hasEdge($v,$w)){return false;}
        $this->g[$v][$w]=true;
        if (!$this->directed){$this->g[$w][$v]=true;}
        $this->m++;
    }
    public function show()
    {
        for ($i=0;$i<$this->n;$i++){
            for ($j=0;$j<$this->n;$j++){
                echo $this->g[$i][$j]."\t";
            }
            echo "\n";
        }
    }
    public function echo()
    {
        return $this->g;
    }
}
//创建一个有3个节点的无向图
$g = new DenseGraph(3,false);
//给节点0和1添加一条边
$g->addEdge(0,1);
//打印图
print_r($g->echo());
//稀疏图-邻接表
//0 1 2   0->1,0->2              0
//1 2     1->2                   |---
//2                              |   |
//			         1---2
class SparseGraph{
    private $n,$m,$directed,$g=[];
    public function __construct($n,$directed)
    {
        $this->n=$n;//定点个数
        $this->m=0;//初始边都为0
        $this->directed=$directed;//是否有向(是a->b还是a<->b)
        for ($i=0;$i<$n;$i++){
            array_push($this->g,[]);
        }
    }
    public function V(){return $this->n;}
    public function E(){return $this->m;}
    private function hasEdge($v,$w){
        if ($v<0 || $v>=$this->n){return false;}
        if ($w<0 || $w>=$this->n){return false;}
        for ($i=0;$i<count($this->g[$v]);$i++){
            if($this->g[$v][$i]==$w)return true;
            return false;
        }
    }
    public function addEdge($v,$w)
    {
        if ($v<0 || $v>=$this->n){return false;}
        if ($w<0 || $w>=$this->n){return false;}
        $this->g[$v][]=$w;
        if ($v!=$w && !$this->directed){$this->g[$w][]=$v;}
        $this->m++;
    }
    public function echo()
    {
        return $this->g;
    }
}
//创建一个有3个节点的有向图
$g = new SparseGraph(3,true);
//给0到1添加一条边
$g->addEdge(0,1);
//给0到2添加一条边
$g->addEdge(0,2);
//给1到2添加一条边
$g->addEdge(1,2);
//打印图
var_dump($g->echo());
//在做遍历邻边的时候，邻接矩阵复杂度为o(v)
//邻接矩阵
////0 1 2 3 4 5 6 7 8
//0 0 0 0 1 0 1 0 0 1
//邻接表
//0 3 5 8
//稀疏图的迭代器，遍历出节点的邻边节点
class SparseGraphAdjIterator{
    private $G,$v,$index;
    public function __construct($graph,$v)
    {
        $this->G = $graph;//图
        $this->v = $v;//目标节点
        $this->index = 0;//迭代值
    }
    //返回第一个要迭代的元素
    public function begin()
    {
        $this->index = 0;
        if(count($this->G[$this->v]))return $this->G[$this->v][$this->index];
        return -1;
    }
    //当前迭代的元素的下一个元素
    public function next()
    {
        $this->index ++;
        if($this->index < count($this->G[$this->v]))return $this->G[$this->v][$this->index];
        return -1;
    }
    //判断遍历终止没有
    public function end()
    {
        return $this->index >= count($this->G[$this->v]);
    }
}
function test_s_mock(){
    $N = 20;//节点
    $M = 100;//要添加的邻边数
    $g = new SparseGraph($N,false);
    for ($i=0;$i<$M;$i++){
        $a = rand()%$N;
        $b = rand()%$N;
        $g->addEdge($a,$b);
    }
    for ($v=0;$v<$N;$v++){
        echo $v.":";
        $adj = new SparseGraphAdjIterator($g->echo(),$v);
        for ($w=$adj->begin();!$adj->end();$w=$adj->next()){
            echo $w." ";
        }
        echo "\n";
    }
    echo "\n";
}
test_s_mock();


//地铁站demo
//from,to
$data = [
    ['MUC','LHR'],
    ['JFK','MUC'],
    ['SFO','SJC'],
    ['LHR','SFO'],
];

//寻找从‘JFK’所能到的终点
function findIt($tickets){
    $g=[];
    buildG($tickets,$g);
    $res = [];
    dfs($g,$res,'JFK');
    var_dump($res);
}

function dfs($g,&$res,$point){

    while (isset($g[$point])&&!empty($g[$point])){
        $p = array_shift($g[$point]);
        dfs($g,$res,$p);
    }
    //每次添加元素都从头部添加
    array_unshift($res,$point);
}

function buildG($tickets,&$g){
    foreach ($tickets as $item){
        $from = $item[0];
        $to = $item[1];
        $g[$from][]=$to;
    }
}

findIt($data);



//寻找有向图是否有环demo
$data = [
    ['MUC','LHR'],
    ['JFK','MUC'],
    ['SFO','SJC'],
    ['LHR','SFO'],
    ["SJC",'JFK']
];

function findIt($tickets){
    $g=[];
    //建立图
    buildG($tickets,$g);
    //检测是否有环是通过“染色”
    $color=[];
    //先给每个顶点染色为0，遍历过的就染色为1
    $data = array_keys($g);
    array_map(function ($val)use(&$color){
        $color[$val]=0;
    },$data);

    $res = isCycle($g,$tickets[0][0],$color);
    var_dump($res,$color);
}

function isCycle($g,$point, &$color){
    if (!isset($g[$point])|| empty($g[$point])|| !isset($color[$point]))return false;

    $color[$point]=1;
    foreach ($g[$point] as $next){
        if (!isset($color[$next]))return false;
        if ($color[$next]==1)return true;
        if ($color[$next]==0 && isCycle($g,$next,$color))return true;
    }

    $color[$point]=2;
    return false;
}

function buildG($tickets,&$g){
    foreach ($tickets as $item){
        $from = $item[0];
        $to = $item[1];
        $g[$from][]=$to;
    }
}

findIt($data);


```



## DAG任务调度器-demo
```php
interface Executor{
    public function execute();
}
class Task implements Executor{
    private $id;
    private $name;
    private $state;
    public function __construct($id,$name,$state)
    {
        $this->id = $id;
        $this->name = $name;
        $this->state = $state;
    }
    public function execute()
    {
        echo "Task id:[ {$this->id} ], task name: [{$this->name}] is running\n";
        $this->state = 1;
        return true;
    }
    public function hasExecuted(){return $this->state == 1;}
    public function getId()
    {
        return $this->id;
    }
}
class Digraph {
    private $tasks=[];//['taskId1'=>$taskClass1]
    private $taskPrevMap = [];//['taskId3'=>[$taskId1,$taskId2],'taskId2'=>[$taskId1]]
    public function addEdge($task,$prevTask)
    {
        $taskId = $task->getId();
        $prevTaskId = $prevTask->getId();
        if(!isset($this->tasks[$taskId])||!isset($this->tasks[$prevTaskId]))return false;
        if (!isset($this->taskPrevMap[$taskId])){$this->taskPrevMap[$taskId]=[];}
        if (isset($this->taskPrevMap[$taskId][$prevTaskId]))return false;
        $this->taskPrevMap[$taskId][]=$prevTaskId;
    }
    public function addTask($task)
    {
        $taskId = $task->getId();
        if (isset($this->tasks[$taskId]))return false;
        $this->tasks[$taskId]=$task;
    }
    public function remove($task)
    {
        $taskId = $task->getId();
        if (!isset($this->tasks[$taskId]))return false;
        if (isset($this->taskPrevMap[$taskId]))unset($this->taskPrevMap[$taskId]);
        foreach ($this->taskPrevMap as $taskName=>$taskPrevMap){
            if (in_array($taskId,$taskPrevMap)){unset($taskPrevMap[array_keys($taskPrevMap,$taskId)[0]]);}
        }
    }
    public function getTasks()
    {
        return $this->tasks;
    }
    public function setTasks($tasks)
    {
        $this->tasks = $tasks;
    }
    public function getTaskPrevMap()
    {
        return $this->taskPrevMap;
    }
    public function setTaskPrevMap($taskPrevMap)
    {
        $this->taskPrevMap = $taskPrevMap;
    }
    public function scheduler()
    {
        while (true){
            $todo = [];
            foreach ($this->getTasks() as $taskId=>$task){
                if (!$task->hasExecuted()){
                    if ( isset($this->getTaskPrevMap()[$taskId]) && !empty( $this->getTaskPrevMap()[$taskId] )){
                        $toAdd = true;
                        foreach ($this->getTaskPrevMap()[$taskId] as $prevId){
                            if (!$this->getTasks()[$prevId]->hasExecuted()) {
                                $toAdd = false;
                                break;
                            }
                        }
                        if ($toAdd)$todo[]=$task;
                    }else{
                        $todo[]=$task;
                    }
                }
            }
            if (!empty($todo)){
                foreach ($todo as $task){
                    if (!$task->execute()) {
                        echo "error\n";
                    }
                }
            }else{
                break;
            }
        }
    }
    
    //利用dfs深度遍历方式
    public function schedulerDFS()
    {
        //O(V+E)
        $visited = [];//遍历过
        $list = [];
        $a = array_keys($this->getTasks());
        array_map(function ($v)use(&$visited){
            $visited[$v]=false;
        },$a);
        foreach ($a as $v){
            if(!$visited[$v]) {
                $this->dfs($v,$visited,$list);
            }
        }
        var_dump($list);
    }

    public function dfs($v,&$visited,&$list)
    {
        if (isset($this->getTaskPrevMap()[$v])){
            foreach ($this->getTaskPrevMap()[$v] as $w){
                if(!$visited[$w])$this->dfs($w,$visited,$list);
            }
        }
        $visited[$v]=true;//是否执行
        $list[]=$v;
    }
    
    //最短路径(从s点到t点)
    public function schedulerBFS($s,$t)
    {
        $visited = [];//遍历过
        $pre = [];
        $list = [];
        $a = array_keys($this->getTasks());
        array_map(function ($v)use(&$visited){
            $visited[$v]=false;
            $pre[$v]=-1;
        },$a);


        $this->dbfs($s,$visited,$pre);
        

        if(!$visited[$t]) return [];

        $cur = $t;
        while ($cur!=$s){
            $list[]=$cur;
            $cur = $pre[$cur];
        }

        $list[]=$s;

        return $list;


    }

    //广度优先遍历算法
    public function dbfs($v,&$visited,&$pre)
    {
        $queue=[];
        array_push($queue,$v);
        $visited[$v]=true;
        $pre[$v]=$v;

        while (!empty($queue)){
            $s = array_shift($queue);
            if (isset($this->getTaskPrevMap()[$s])){
                foreach ($this->getTaskPrevMap()[$s] as $w){
                    if(!$visited[$w]){
                        array_push($queue,$w);
                        $visited[$w]=true;
                        $pre[$w]=$s;
                    }
                }
            }
        }
    }
}
$d = new Digraph();
$t1 = new Task(1,'t1',0);
$t2 = new Task(2,'t2',0);
$t3 = new Task(3,'t3',0);
$t4 = new Task(4,'t4',0);
$t5 = new Task(5,'t5',0);
$t6 = new Task(6,'t6',0);
$d->addTask($t3);
$d->addTask($t1);
$d->addTask($t2);
$d->addTask($t6);
$d->addTask($t4);
$d->addTask($t5);
//      3       4
//      |       |
//      |-—>2<—-|     5
//          |         |
//          |         |
//  6<------|--->1<---|      
//               
$d->addEdge($t1, $t2);//先执行t2再执行t1
$d->addEdge($t1, $t5);//先执行t5再执行t1
$d->addEdge($t6, $t2);//先执行t2再执行t6
$d->addEdge($t2, $t3);//先执行t3再执行t2
$d->addEdge($t2, $t4);//先执行t4再执行t2
$d->scheduler();
//输出的执行顺序为：
//Task id:[ 3 ], task name: [t3] is running
//Task id:[ 4 ], task name: [t4] is running
//Task id:[ 5 ], task name: [t5] is running
//Task id:[ 2 ], task name: [t2] is running
//Task id:[ 1 ], task name: [t1] is running
//Task id:[ 6 ], task name: [t6] is running
```
</br></br>
```golang
    type Task struct{
	Name string
	Commend string
	PrevTask map[string]*Task
	Status bool
}

type WorkFlow struct {
	task map[string]*Task
}

func(wf *WorkFlow)add(t *Task){
	if _, ok := wf.task[t.Name]; ok {
		fmt.Println("err,exsit")
	}
	wf.task[t.Name]=t

}

func(wf *WorkFlow)run(){
	for _,t := range wf.task{
		if !t.Status{
			wf.dsf(t)
		}
	}
}

//深度优先遍历
func (wf *WorkFlow)dsf(t *Task)  {
	if t.PrevTask != nil {
		for _,pt := range t.PrevTask{
			if !pt.Status{wf.dsf(pt)}
		}
	}
	fmt.Println(t.Name,t.Commend)
	t.Status=true
}

func NewWorkFlow()*WorkFlow{
	return &WorkFlow{
		task:make(map[string]*Task),
	}
}
func main() {

	wf := NewWorkFlow()
	t1:=&Task{
		Name:"t1",
		Commend:"echo 1",
		PrevTask:make(map[string]*Task),
		Status:false,
	}
	t2:=&Task{
		Name:"t2",
		Commend:"echo 2",
		PrevTask:make(map[string]*Task),
		Status:false,
	}
	t3:=&Task{
		Name:"t3",
		Commend:"echo 3",
		PrevTask:make(map[string]*Task),
		Status:false,
	}

	wf.add(t1)
	wf.add(t3)
	wf.add(t2)

	t1.PrevTask[t2.Name]=t2
	t2.PrevTask[t3.Name]=t3

	wf.run()
}
    
```
</br></br>
```golang
//优化版
/任务抽象
type WorkFlowTask interface {
	Name()string
	Ready() bool
	Run()
}

//任务实例
type Task struct {
	name string
	Commend string
	PrevTask map[string]*Task
	status bool
}
func(t *Task)Name()string{
	return t.name
}
func(t *Task)Ready()bool{
	return t.status
}
//深度优先遍历(dfs)
func(t *Task)Run(){
	if t.PrevTask != nil{
		for _,pt:=range t.PrevTask{
			if !pt.Ready(){pt.Run()}
		}
	}
	fmt.Println(t.name,t.Commend)
	t.status=true
}


//引擎
type WorkFlow struct {
	task map[string]WorkFlowTask
}
func(wf *WorkFlow)add(t WorkFlowTask){
	if _, ok := wf.task[t.Name()]; ok {
		fmt.Println("err,exsit")
	}
	wf.task[t.Name()]=t
}
func(wf *WorkFlow)run(){
	for _,t := range wf.task{
		if !t.Ready(){
			t.Run()
		}
	}
}
func NewWorkFlow()*WorkFlow{
	return &WorkFlow{
		task:make(map[string]WorkFlowTask),
	}
}

func main() {

	wf := NewWorkFlow()
	t1:=&Task{
			name:"t1",
			Commend:"echo 1",
			PrevTask:make(map[string]*Task),
			status:false,
	}
	t2:=&Task{
		name:"t2",
		Commend:"echo 2",
		PrevTask:make(map[string]*Task),
		status:false,
	}
	wf.add(t1)
	wf.add(t2)
	t1.PrevTask[t2.Name()]=t2

	wf.run()
}
```

## 图在树形层级应用
```php
class categoryParse{
    //存储子类目与父类目映射
    public $graph=[];
    //存储id和name的映射
    public $map=[];
    public function __construct($json)
    {
        $data = json_decode($json,true);
        //建图
        $this->traverse($data,0);
    }

    private function traverse($data,$pre)
    {
        foreach ($data as $value){
            $pre?
                $this->graph[$value['id']][]=$pre:
                $this->graph[$value['id']]=[];

            $this->map[$value['id']]=$value['name'];
            if (array_key_exists('children',$value)){
                $this->traverse($value['children'],$value['id']);
            }

        }
    }

    /**
     *  遍历子类目到总父类目的关系链
     * @param $id integer 类目id
     * @return array 关系链
     */
    private function DFS($id){
        if (!isset($this->graph[$id])){
            return [];
        }
        $stack=new SplStack();
        $stack->push($this->map[$id]);
        $ele=$id;
        while (!empty($this->graph[$ele])){
            $ele=$this->graph[$ele][0];
            $stack->push($this->map[$ele]);
        }
        $array=[];
        foreach ($stack as $item){
            $array[]=$item;
        }

        return $array;
    }

    //传入指定类目id，打印出从根类目到指定类目的关系链
    public function printCategory($id){
        $resArr = $this->DFS($id);
        echo empty($resArr)?
            "类目不存在".PHP_EOL:
            implode(" > ",$resArr).PHP_EOL;
    }

    //生成无限层级json结构
    public function treeData()
    {
        $tree =[];
        $tmpData = []; //临时数据,保存节点结构
        $point = array_keys($this->map);

        foreach ($point as $item){
            $tmpData[$item]=[
                'id'=>$item,
                'name'=>$this->map[$item]
            ];
        }
        foreach ($this->graph as $item=>$pre){
            if (!empty($pre)&&isset( $tmpData[$pre[0]])){
                $tmpData[$pre[0]]['children'][] = &$tmpData[$item];
            }else{
                //一级分类
                $tree[] = &$tmpData[$item];
            }
        }

        unset($tmpData);
        return json_encode($tree,JSON_UNESCAPED_UNICODE);

    }

}

$json = file_get_contents("https://job.xiyanghui.com/api/q1/json");
$categoryParse = new categoryParse($json);
$categoryParse->printCategory(2221);
echo $categoryParse->treeData();
```


## 最短路径
```php
class Graph{
    public $nodes=[];
    public $graph = [];

    public function addNode($node)
    {
        $this->nodes[]=$node;
    }
    
    public function addWeightEdge($node,$to,$weight)
    {
        $edge = new stdClass();
        $edge->to=$to;
        $edge->weight=$weight;
        $this->graph[$node][]=$edge;
    }
    
    private function dag($start)
    {
        $visited = [];//遍历过
        $prev=[];//存储最短路径过程
        $dist = [$start=>0];
        array_map(function ($v)use(&$visited){
            $visited[$v]=false;
        },$this->nodes);
        $pg = new SplQueue();
        $pg->enqueue([$start,0]);

        while (!$pg->isEmpty()){
            list($index,$minValue) = $pg->dequeue();
            $visited[$index]=true;

            //如果之前加入队列的点最小值被后来的点路径更改，权重图已经是最小则跳过
            if ($dist[$index]<$minValue)continue;

            //遍历邻接边
            if (isset($this->graph[$index])){
                $edge = $this->graph[$index];
                foreach ($edge as $item){
                    //访问过就跳过
                    if ($visited[$item->to])continue;
                    //当前点到邻接点的权重为当前权重+到邻接点的权重
                    $newDist = $dist[$index]+$item->weight;

                    //到邻接点要取最小的值路径保留
                    if (!isset($dist[$item->to])){
                        $dist[$item->to]=$newDist;
                        $prev[$item->to]=$index;
                    }
                    else {
                        if ($dist[$item->to]>$newDist){
                            $dist[$item->to]=$newDist;
                            $prev[$item->to]=$index;
                        }
                    }

                    $pg->enqueue([$item->to,$dist[$item->to]]);
                }
            }
        }
        return [$dist,$prev];

    }

    public function findShortestPath($s,$e)
    {
        list($dist,$prev)=$this->dag($s);
        $path = new SplStack();
        if (!isset($dist[$e]))return $path;
        for($at=$e;isset($prev[$at]);$at=$prev[$at]){
            $path->push($at);
        }
        $path->push($s);
        return $path;
    }
}

$g = new Graph();
for ($i = 0; $i < 7; $i++) {
    $g->addNode($i);
}

$g->addWeightEdge(0,1,3);
$g->addWeightEdge(0, 2, 2);
$g->addWeightEdge(0, 5, 3);
$g->addWeightEdge(1, 3, 1);
$g->addWeightEdge(1, 2, 6);
$g->addWeightEdge(2,3,1);
$g->addWeightEdge(2,4,10);
$g->addWeightEdge(3,4,5);
$g->addWeightEdge(5,4,7);

$a = $g->findShortestPath(0,4);
foreach ($a as $item){
    echo $item.PHP_EOL;//0->2->3->4
}

```


## 获取数据流中的中位数
```php
//原理：维护2个堆，一个大顶堆一个小顶堆
//小顶堆元素都比大顶堆的元素大
//2个堆的堆顶元素是中位数候补或者构成中位数的集合
//元素的插入根据上一次的顺序来，第一个元素放入小顶堆，第2个元素放入大顶堆，以此交替放入
//当然元素放入前要先跟堆顶元素做比较：如果要放入小顶堆的，若插入元素比大顶堆元素小，则要把大顶堆堆顶元素弹出放入小堆顶，而插入元素放入大顶堆
//放入大顶堆的情况也一样
//法则是：小顶堆元素都比大顶堆的元素大

class M{
    private $minHeap;
    private $maxHeap;
    private $count=0;

    public function __construct()
    {
        $this->minHeap = new SplMinHeap();
        $this->maxHeap = new SplMaxHeap();
    }

    public function Insert($num)
    {
        //【偶数小顶堆】
        if ($this->count%2==0){
            //如果插入的数字比大堆顶元素小
            if (!$this->maxHeap->isEmpty()&&$this->maxHeap->top()>$num){
                $oldMax = $this->maxHeap->extract();
                $this->maxHeap->insert($num);
                $num=$oldMax;
            }
            $this->minHeap->insert($num);
        }
        //【奇数大顶堆】
        else{
            //如果插入的数字比小堆顶元素大
            if (!$this->minHeap->isEmpty()&&$this->minHeap->top()<$num){
                $oldMin = $this->minHeap->extract();
                $this->minHeap->insert($num);
                $num=$oldMin;
            }
            $this->maxHeap->insert($num);
        }
        $this->count++;
    }

    public function getMedian()
    {
        //统计数据大小
        $size = $this->minHeap->count() + $this->maxHeap->count();
        if($size==0)return 0.0;

        //如果数据为偶数
        if(($size & 1) == 0)return ($this->minHeap->top() + $this->maxHeap->top()) / 2.0;

        //如果为奇数
        return doubleval($this->minHeap->top());

    }
}

$m = new M();
$m->Insert(7);
$m->Insert(5);
$m->Insert(4);
$m->Insert(9);
var_dump($m->getMedian());//6
```


## 逆波兰式使用场景
```php
var_dump('test');
```


## GeoHash算法在LBS的应用
案例1：当打开外卖APP,搜索附近美食,数据总可以按照由近到远展示在面前。
</br></br>
案例2：打开信息流APP,可以看到附近5公里范围内最新的Feed。
</br></br>
上面的案例都是使用到了`GeoHash`算法：GeoHash将二维的经纬度转换成字符串，如坐标(39.92324, 116.3906)就可以转换成wx4g0ec1
</br></br>
案例一的实现思路：首次获取用户的经纬度信息，通过地图API(如百度)获取附近商圈的信息(由近到远)，再将经纬度信息经过GeoHash获取字符串编码，弄一缓存来存储该信息，以字符串编码为key，商圈信息为value；下次如果有相似字符串编码(前缀匹配)可直接在缓存中查询。
</br></br>
案例二的实现思路：根据Feed的经纬度，计算geohash，空间索引使用Redis的zset结构，将geohash作为空间索引的key，feed_id作为member，时间作为score。
查询时根据用户当前经纬度，计算geohash，就能找到他附近的Feed。但存在边界问题，附近的Feed不一定在同一个矩形区域内。解决这个问题可以在查询时扩大范围，除了查询用户所在的矩形外，还扩散搜索相邻的8个矩形，将9个矩形合并（如下图），按时间排序，过滤掉超出距离范围的Feed，最后做分页查询。但是这种方式查询效率比较低，作为读远远大于写的场景，换了一种思路，在更新Feed空间索引时，将Feed写入相邻的8个矩形，这样每个矩形还包含了相邻矩形的Feed。通过数据冗余的方式，换取了更高的查询效率。将复杂的geo查询，简化为redis的zrange操作，性能提高了一个数量级，平均耗时降到3ms。空间索引通过geohash分片到redis节点，具有数据分布均匀、方便扩容的优势。
</br></br>
geohash实现原理：[https://www.cnblogs.com/dengxinglin/archive/2012/12/14/2817761.html]
</br></br>
参考案例：[https://www.codetd.com/article/4267648]和[https://toutiao.io/posts/i62kms/preview]





## MySQL的in的查询结果集按顺序
MySQL in 查询，并通过 FIELD 函数按照查询条件顺序返回结果，详细概念可以参考文章[http://martin91.github.io/blog/articles/2015/09/13/mysql-in-query-and-order-by-field-function/] 

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
### 在win下编译linux可执行文件
cmd控制台到要编译的go文件的目录下
<br>
键入以下命令：【GOOS:目标平台的操作系统(linux、darwin、freebsd、window)；GOARCH:目标平台的体系结构(386[编译32位可执行程序]、amd64
、arm)；交叉编译不支持CGO所以禁用它】
```shell
>set CGO_ENABLED=0
>set GOOS=linux
>set GOARCH=amd64
>go build main.go
```
即可生成一个没有后缀的二进制文件，然后将该文件放到linux系统某个文件下，并赋予权限
```shell
>chmod 777 main
```
最后就可以执行了:
```shell
>./main
```
如果想让文件在后台执行,命令为：
```shell
>nohub ./main &
```


### defer执行顺序
多个defer的执行顺序是以栈的特性先进后出来执行，越写在前面就越后执行

### 初始化二维数组
```golang
    x := 2
    y := 4

    table := make([][]int,x)
    for i:= range table {
        table[i] = make([]int,y)
    }
    
```


### 没有前置运算符
不存在++a或者--a，只有a++和a--


### struct tag
只要在相关tag中写上`omitempty`，只要在struct初始化时没有赋值的字段，就会省略掉，如：
```golang
type Response struct{
    Code int `json:"code" bson:"name,omitempty"`
    Msg  string `json:"msg,omitempty"`
}
r1 := &Response{Code:404}
r2 := &Response{Msg:"error"}
```

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


### httpApi请求
```go
package main
import (
    "io/ioutil"
    "net/http"
    "net/url"
    "fmt"
    "encoding/json"
)

//1.全部开放基金
func Request1()[]byte{
    //请求地址
    juheURL :="http://web.juhe.cn:8080/fund/netdata/all"
 
    //初始化参数
    param:=url.Values{}
 
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("key",APPKEY) //APPKEY值
 
 
    //发送请求
    data,err:=Get(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
	return nil
    }
    
    return data
}
 
 
// get 网络请求
func Get(apiURL string,params url.Values)(rs[]byte ,err error){
    var Url *url.URL
    Url,err=url.Parse(apiURL)
    if err!=nil{
        fmt.Printf("解析url错误:\r\n%v",err)
        return nil,err
    }
    //如果参数中有中文参数,这个方法会进行URLEncode
    Url.RawQuery=params.Encode()
    resp,err:=http.Get(Url.String())
    if err!=nil{
        fmt.Println("err:",err)
        return nil,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
 
// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values)(rs[]byte,err error){
    resp,err:=http.PostForm(apiURL, params)
    if err!=nil{
        return nil ,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}

func main(){
 
 
 //json数据：{
    "error_code": 0,
    "reason": "success",
    "total": 601,
    "result": [
        {
            "jjlx": "偏股型基金",
            "nav_a": "-0.004",
            "nav_date": "2015-12-10",
            "nav_rate": "-0.2959",
            "per_nav": "1.3480",
            "sg_states": "开放",
            "sname": "信诚中证500指数",
            "symbol": "165511",
            "total_nav": "1.8300",
            "yesterday_nav": "1.352"
        }
    ]
}
    var netReturn map[string]interface{}
    //1.全部开放基金
    data:=Request1()
    
    json.Unmarshal(data,&netReturn)
    if netReturn["error_code"].(float64)==0{
       fmt.Printf("接口返回result字段是:\r\n%v",netReturn["result"])
       //要获取里面的nav_rate
       a:=netReturn["result"].([]interface{})
       for _,v:= range a{
       fmt.Println(v.(map[string]interface{})["total_nav"])
       }
    }
 
}

```



### 高效字符串拼接方法
```go
func main(){
	//字符串拼接
	//1.10以後的版本
	var builder strings.Builder
	for i:=0;i<10;i++{
		builder.WriteString("1")
	}
	fmt.Println(builder.String())
	
	
	//以前的版本
	var buf bytes.Buffer
	for i:=0;i<10;i++{
		buf.WriteString("1")
	}
	fmt.Println(buf.String())
}
```

### 生成随机数
```go
import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	//以时间戳作为随机数种子
	seed := time.now().UnixNano()
	//生成一个0到9的随机数
	code:= rand.New(rand.NewSource(seed)),Intn(10)
	
	//随机获取切片的一个元素
	List = []string{"m","a","r","c","u","s"}
	count := len(List)
	index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
	elem := List[index]
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

### 利用sync.Once构建单例
```go
type Singleton struct {}
var (
    singleInstance *Singleton
    once sync.Once
    )
func GetSingletonObj() *Singleton {
	//只会执行一次
   once.Do(func() {
	   fmt.Println("Create obj")
	   singleInstance = new(Singleton)
   })
   return singleInstance
}
func main(){
   var wg sync.WaitGroup
   for i:=0;i<10 ;i++  {
	wg.Add(1)
	go func() {
	   obj := GetSingletonObj()
	   fmt.Printf("%x\n",unsafe.Pointer(obj))//打印内存地址
	   wg.Done()
	}()
    }
    wg.Wait()
}
```

### 利用buffer Channel来解决开多任务协程泄漏问题
```go
//案例情景：开多协程来执行任务,只要一收到结果马上停止程序
//优化前：
//任务函数
func runTask(i int) string {
	//Do something …………
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d",i)
}
//
func getResponse() string {
	//协程数
	taskNum := 10
	//准备一个channel来接受协程返回的数据
	c := make(chan string)
	for i:=0;i< taskNum;i++{
		go func(i int) {
			ret := runTask(i)
			c <- ret
		}(i)
	}
	//获取数据
	return <-c
}
func main(){
   	//打印函数执行前协程数
	fmt.Println("Before:",runtime.NumGoroutine())
	fmt.Println("receve:",getResponse())
	time.Sleep(1*time.Second)
	//打印函数执行后的协程数
	fmt.Println("after:",runtime.NumGoroutine())
}
//上面的程序执行后我们注意到"after:"打印出的协程数非常多，证明在getResponse()函数执行完后，之前开的协程并没有结束还在内存中
//造成的原因是因为创建的channel不是buffer的，这样就会造成只有最快的一个协程可以把数据放入，其他只会一直阻塞在那里，协程无法释放
//优化方案：把c := make(chan string)改成c := make(chan string,taskNum)即可，协程们只要把数据放进去就可以立刻释放
```

### 利用close(chan)来停止goroutine
```go
//判断channel是否已经关闭
func isCancelled(cancelChan <-chan struct{}) bool  {
	select {
		case <-cancelChan: return true
		default: return false
	}
}
func cancel_1(cancelChan chan struct{})  {
	cancelChan <- struct{}{}
}
//广播式
func cancel_2(cancelChan chan struct{})  {
	close(cancelChan)
}
func testCancel()  {
	cancelChan := make(chan struct{},0)
	for i:=0;i<5;i++{
		go func(i int,cancelCh chan struct{}) {
			for  {
				if isCancelled(cancelCh){break}
				time.Sleep(5*time.Millisecond)
			}
			fmt.Println(i,"Cancelled")
		}(i,cancelChan)
	}
	//cancel_1(cancelChan)//只会打印 0 Cancelled，因为只有最快的一个可以收到关闭信息
	cancel_2(cancelChan)//会打印全部
	time.Sleep(1*time.Second)
}
```

### 利用buffer Channel来构成对象池
```go
//对象类型
type ReusableObj struct {}
//对象池
type ObjPool struct {
	bufChan chan *ReusableObj //用于缓冲可重用的对象
}
//创建对象池
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj,numOfObj)
	for i:=0;i<numOfObj;i++{
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}
//获取对象
func (p *ObjPool)GetObj(timeout time.Duration) (*ReusableObj,error)  {
	select {
	case ret:= <-p.bufChan:
		return ret,nil
	case <-time.After(timeout):
		//超时控制
		return nil,errors.New("time out")
	}
}
//释放对象
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan<- obj:
		return nil
	default://当放入bufChan发生异常(如chan满了或者放入类型不对)就会走这个分支
		return errors.New("overflow")
	}
}
func main(){
	pool := NewObjPool(10)
	//释放
	if err:= pool.ReleaseObj(&ReusableObj{});err!=nil{
		fmt.Println(err)
	}
	for i:=0;i<11;i++{
		if v,err:= pool.GetObj(3*time.Second);err!=nil{
			fmt.Println(err)
		}else {
			fmt.Println("%T\n",v)
			if err := pool.ReleaseObj(v);err!=nil{
				fmt.Println(err)
			}
		}
	}
}
```

### sync.Pool对象缓存
```go
//sync.Pool对象的生命周期
//GC会清除其缓存的对象
//对象的缓存有效期为下一系的GC之前
pool := &sync.Pool{
	New: func() interface{} {
		fmt.Println("Create a new object");
		return 100
	},
}
v := pool.Get().(int)
fmt.Println(v)//100
pool.Put(3)
runtime.GC()//GC 会清除sync.Pool中的缓存对象
v1,_:=pool.Get().(int)
fmt.Println(v1)//100
//适合于通过复用来降低复杂对象的创建和GC的代价
//协程安全，但会有锁的开销
//生命周期受GC的影响，不适合于做连接池等，需要自己管理生命周期的资源的池化
```

### ctrl+c退出程序
```go
func mian(){
	go func(){}()
	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
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


### 事件分发组件模式，实现了常规的广播，队列等事件分发模型
```go
import (
	"errors"
	"sync"
)

type Event interface {}

var ErrSinkClosed = errors.New("ErrSinkClosed")

//Sink是一个用来分发事件（Event）的结构。可以当作事件的处理者，使用接口的方式声明。
//只要对象实现了这两个方法，就可以被当作一个Sink。
type Sink interface {
	Write(event Event) error
	Close() error
}


type configureRequest struct {
	sink Sink
	response chan error
}


//Boardcaster由多个Sink组成
type Broadcaster struct {
	sinks   []Sink //所包含的Sink
	events  chan Event// 同步Event的channel
	adds    chan configureRequest //adds和remove必须保证thread-safe，所以采用channel同步
	removes chan configureRequest

	shutdown chan struct{}
	closed   chan struct{}
	once     sync.Once
}

//在Broadcaster中所有的临界资源(sinks,event)都通过自身的run()函数统一管理，外界则通过相应的channel 同步给Broadcaster
//例如Write()
func (b *Broadcaster) Write(event Event) error {
	select {
	case b.events <- event:
	case <-b.closed:
		return ErrSinkClosed
	}
	return nil
}


//可以看到增减sink都是通过向对应的channel写入数据进行的。
func (b *Broadcaster) Add(sink Sink) error {
	return b.configure(b.adds, sink) //  will be block until ch can be writen
}

func (b *Broadcaster) configure(ch chan configureRequest, sink Sink) error {
	response := make(chan error, 1)

	for {
		select {
		case ch <- configureRequest{
			sink:     sink,
			response: response}:
			ch = nil // ？
		case err := <-response:
			return err
		case <-b.closed:
			return ErrSinkClosed
		}
	}
}
//核心run函数的实现,监听Boardcast管道上的相应事件，并作出处理。
func (b *Broadcaster) run() {
	defer close(b.closed)
	//将remove封装了一下，因为下面两处都会用到
	remove := func(target Sink) {
		for i, sink := range b.sinks {
			if sink == target {
				b.sinks = append(b.sinks[:i], b.sinks[i+1:]...)
				break
			}
		}
	}
	// 轮训处理channel上的事件
	for {
		select {
		case event := <-b.events: //有事件到来，进行广播
			for _, sink := range b.sinks {
				if err := sink.Write(event); err != nil {
					if err == ErrSinkClosed {
						// remove closed sinks
						remove(sink)
						continue
					}
					//logrus.WithField("event", event).WithField("events.sink", sink).WithError(err). Errorf("broadcaster: dropping event")
				}
			}
		case request := <-b.adds: //增加sink事件
			// while we have to iterate for add/remove, common iteration for
			// send is faster against slice.

			var found bool
			for _, sink := range b.sinks {
				if request.sink == sink {
					found = true
					break
				}
			}

			if !found {
				b.sinks = append(b.sinks, request.sink)
			}
			// b.sinks[request.sink] = struct{}{}
			request.response <- nil // 唤醒阻塞的configure（）函数

		case request := <-b.removes://删除sink事件
			remove(request.sink)
			request.response <- nil
		case <-b.shutdown:
			// close all the underlying sinks
			for _, sink := range b.sinks {
				if err := sink.Close(); err != nil && err != ErrSinkClosed {
					//logrus.WithField("events.sink", sink).WithError(err).
					//	Errorf("broadcaster: closing sink failed")
				}
			}
			return
		}
	}
}
```


### 反射
```go
//reflect.TypeOf() 返回类型(reflect.Type)
//reflect.ValueOf() 返回值(reflect.Value)
//可以从reflect.Value中获得类型
//通过kind来判断类型
func CheckType(v interface{})  {
	t:= reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32,reflect.Float64:
		fmt.Println("Float")
	case reflect.Int,reflect.Int32,reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown",t)
	}
}
func testBasicType()  {
	var f float64 = 12
	CheckType(f)//Float
	CheckType(&f)//Unknown &float64
}
func testTypeAndValue()  {
	var f int64 = 10
	fmt.Println(reflect.TypeOf(f),reflect.ValueOf(f))//int64,10
	fmt.Println(reflect.ValueOf(f).Type())//int64
}
//利用反射编写灵活的代码
//按名字访问结构的成员
//reflect.ValueOf(*e).FieldByName("name")
//按名字访问结构的方法
//reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
//注意!:reflect.Type和reflect.Value都有FieldByName方法，但却有区别的
type Employee struct {
	Employee string
	Name string `format:"normal"`
	Age int
}
func (e *Employee) UpdateAge(newValue int)  {
	e.Age = newValue
}
type Customer struct {
	Cookie string
	Name string
	Age int
}
func testInvokeByName()  {
	e:=&Employee{"1","Mike",30}
	//按名字获取成员
	fmt.Printf("Name: value(%[1]v),Type(%[1]T)",
		reflect.ValueOf(*e).FieldByName("Name"))
	//Name: value(Mike),Type(reflect.Value)
	if nameField,ok:= reflect.TypeOf(*e).FieldByName("Name");!ok {
		//不存在该属性
		fmt.Println("Failed to get 'Name' field")
	}else {
		//访问structTag
		fmt.Println("Tag:format ",
			nameField.Tag.Get("format"))
		//Tag:format normal
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)})
	fmt.Println("Updated Age:",e)// Updated Age:&{1 Mike 1}
}
```

### log日志收集到一个文件
```go
func main() {
	// 按照所需读写权限创建文件
	f, err := os.OpenFile("filename.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// 完成后延迟关闭，而不是习惯!
	defer f.Close()
	//设置日志输出到 f
	log.SetOutput(f)
	//测试用例
	log.Println("check to make sure it works")
}
```
运行后filename.log文件的内容为：2019/09/20 10:20:34 check to make sure it works


### 利用reflect.DeepEqual来让切片之间或者map之间进行比较判断
```go
//DeepEqual
	//切片与map的比较(原本是无法进行比较的)
	a := map[int]string{1:"one",2:"two",3:"three"}
	b := map[int]string{1:"one",2:"two",4:"four"}
	//直接a == b 是会报语法错误的
	//而用reflect.DeepEqual就可以来判断
	fmt.Println(reflect.DeepEqual(a,b))
	//切片也可以比较了
	s1 := []int{1,2,3}
	s2 := []int{1,2,3}
	fmt.Println(reflect.DeepEqual(s1,s2))
```


### 实时读取显示文件最新内容
```go
type LogProcess struct {
	rc chan []byte //读取器把读取的数据放入rc
	wc chan string //把从rc读取到的数据巾帼处理放入到wc
	reader Reader //读取器
	writer Writer //写入器
}
func (l *LogProcess)Process()  {
	//解析模块
	for v := range l.rc{
		l.wc <- strings.ToUpper(string(v))
	}
}
type Reader interface {
	read(rc chan []byte)
}
type Writer interface {
	write(wc chan string)
}
type txtRead struct {
	path string
}
func (r *txtRead)read(rc chan []byte)  {
	f,err := os.Open(r.path)
	if err!=nil{
		panic(fmt.Sprintf("open file error:%s",err.Error()))
	}
	//从文件末尾开始逐行读取文件内容
	f.Seek(0,2)
	rd := bufio.NewReader(f)
	for  {
		line,err := rd.ReadBytes('\n')
		if err == io.EOF{
			time.Sleep(500 * time.Millisecond)
			continue
		}else if err!=nil {
			panic(fmt.Sprintf("ReadBytes error:%s",err.Error()))
		}
		//把换行符去掉
		rc <- line[:len(line)-1]
	}
}
type txtWrite struct {}
func (r *txtWrite)write(wc chan string)  {
	for v := range wc{
		fmt.Println(v)
	}
}
func main(){
	r := &txtRead{path:"test.log"}
	w := &txtWrite{}
	l := &LogProcess{
		rc:make(chan []byte),
		wc:make(chan string),
		reader:r,
		writer:w,
	}
	go l.reader.read(l.rc)
	go l.Process()
	go l.writer.write(l.wc)
	select{}
}
```

```shell
echo hello >> test.log
echo marcus >> test.log
```

### 扇入模式
```go
func testFanIn(){
	//假设有2个channel来接受2个协程的数据
	c1,c2:=make(chan int),make(chan int)
	//现在要将2个channel的数据合并成一个channel
	//需要设计这样的扇入函数
	merge:= func(cs...<-chan int) <-chan int {
		//利用wg来控制同步
		var wg sync.WaitGroup
		//创建一个统一的输出的channel
		out := make(chan int)
		//为每个输入 channel 启动一个 goroutine
		output := func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}
		//控制同步
		wg.Add(len(cs))
		//开始合并数据
		for _,c := range cs{
			go output(c)
		}
		// 启动一个 goroutine 负责在所有的输入 channel 关闭后，关闭这个唯一的输出 channel
		go func() {
			wg.Wait()
			close(out)
		}()
		return out
	}
	//将数据统一取出
	for n := range merge(c1,c2){
		fmt.Println(n)
	}
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


### 任务调度器专栏
#### 利用os/exec包来执行linux命令
```go
import (
   "os/exec"
   "log"
   "error"
   "fmt"
   "context"
   "time"
)
//在win10中的用户文件夹中存在一个隐藏的文件夹与文件/bin/bash.exe
//利用它可以向linux一样执行linux命令
func demo1(){
   //配置要执行的命令:参数1为bash器，后面的参数可自由添加
   //此处的命令为:列出当前文件，然后沉睡1s，往test.log文件追加内容"golang"；然后沉睡5s，输出ok
   cmd := exec.Command("./bash.exe","-c",
	"ls -l;sleep 1;echo golang >> test.log;sleep 5;echo ok")
   //執行命令，并且去捕获子进程的输出(pipe)
   output, err := cmd.CombinedOutput()
   if err!=nil{
	log.Fatal("error:",err)
	return
   }
   //打印子进程的输出(output为[]byte)
   fmt.Println(string(output))
}
//用于存放命令执行后的结果信息体
type Res struct {
   output []byte
   err error
}
func demo2(){
	//执行一个cmd，让其在一个协程中执行，让其执行2s;sleep 2;echo hello;
	//1s的时候我们杀死进程,然后看看会结果返回什么
	var (
		ctx context.Context//chan byte
		cancelFunc context.CancelFunc//close(chan byte)
		cmd *exec.Cmd
		resChan chan *Res
		res *Res
	)
	//创建一个上下文
	ctx , cancelFunc = context.WithCancel(context.TODO())
	//创建一个结果队列
	resChan = make(chan *Res,1000)
	//创建子协程
	go func() {
		var  (
			output []byte
			err error
		)
		//配置命令
		cmd = exec.CommandContext(ctx,"./bash.exe","-c","sleep 2;echo hello")
		//执行任务，捕获输出
		output,err = cmd.CombinedOutput()
		//把人物输出结果传回main协程
		resChan<- &Res{output:output,err:err}
	}()
	//沉睡1s
	time.Sleep(1*time.Second)
	//取消上下文
	cancelFunc()
	//等待子协程退出，并打印任务执行结果
	res = <- resChan
	//打印结果信息
	fmt.Println(res.err,string(res.output))
	//打印出：exit status 1
}
```

#### 利用gorhill/cronexpr包与cron命令来构建定时任务器
```go
//Cron常见用法
//  * * * * *
//  分(0-59) 时(0-23) 日(1-31) 月(1-12) 星期(0-7)
//  */5 * * * * echo hello > test.log 每5分钟执行一次
//  1-5 * * * * echo hello > test.log 第1-5分钟执行5次
//  0 10,20 * * * echo bye | tail -1 每天10点,22点整执行一次
//使用第三方库来识别cron命令
//go get github.com/gorhill/cronexpr
type CronJob struct {
	expr *cronexpr.Expression
	nextTime time.Time
}
func demo3()  {
	var (
		expr *cronexpr.Expression
		err error
		now,nextTime time.Time
	)
	//* * * * * * *
	//秒粒度，年配置(2018-2099)
	//每5秒执行1次
	if expr,err = cronexpr.Parse("*/5 * * * * * *");err!=nil{
		fmt.Println("error:",err)
		return
	}
	//当前时间
	now = time.Now()
	//下次调度时间
	nextTime = expr.Next(now)
	//等待这个定时器超时
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("下次被调度时间：",nextTime)
	})
	time.Sleep(5*time.Second)
}
type CronJob struct {
	//cron表达式
	expr *cronexpr.Expression
	//下一次的执行时间
	nextTime time.Time
}
func demo4(){
	//需要有1个调度协程，它能定时检查所有的cron任务，谁过期了就执行谁
	var (
		scheduleTable map[string]*CronJob
		now time.Time
		expr *cronexpr.Expression
		cron *CronJob
	)
	//存儲表
	scheduleTable = make(map[string]*CronJob)
	//当前时间
	now = time.Now()
	//定义2个cronJob
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cron = &CronJob{
		expr:expr,
		nextTime:expr.Next(now),
	}
	scheduleTable["job1"] = cron
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cron = &CronJob{
		expr:expr,
		nextTime:expr.Next(now),
	}
	scheduleTable["job2"] = cron
	//启动一个调度协程
	go func() {
		//定时检查一下任务
		for  {
			now := time.Now()
			for k,v := range scheduleTable{
				//判断是否过期
				if v.nextTime.Before(now) || v.nextTime.Equal(now){
				   //启动协程执行这个任务
				   go func(name string) {
					fmt.Println("执行任务",name)
				   }(k)
				   //计算下一次调度的时间
				   v.nextTime = v.expr.Next(now)
				   fmt.Println(k,"下次任务执行的时间:",v.nextTime)
				}
			}
			select {
			   //睡眠100毫秒后返回
			   case <-time.NewTimer(100*time.Millisecond).C:
			}
		}
	}()
	//睡眠100s观察状况
	time.Sleep(100*time.Second)
}
```

#### etcd入门操作
详细相关可以参考文章[https://yuerblog.cc/2017/12/12/etcd-v3-sdk-usage/] 
```go
//下载官方的库：go get go.etcd.io/etcd/clientv3
//初始化etcd与创建kv
func initAndPutKV()  {
   var (
	config clientv3.Config
	client *clientv3.Client
	err error
	kv clientv3.KV
	putResp *clientv3.PutResponse
   )
   //初始化 start-------
   //客户端配置
   config = clientv3.Config{
	Endpoints:[]string{"127.0.0.1:2379"},
	DialTimeout:5*time.Second,
   }
   //建立连接
   if client,err = clientv3.New(config);err!=nil{
	fmt.Println(err)
	return
   }
   //初始化 end-------
   //用于读写etcd的kv键值对
   kv = clientv3.NewKV(client)
   //添加k（clientv3.WithPrevKV()为查看前一次的v值）
   if putResp,err = kv.Put(context.TODO(),"/test/job1","is best",clientv3.WithPrevKV());err!=nil{
	fmt.Println(err)
   }else{
	//查看当前版本
	fmt.Println("Revision:",putResp.Header.Revision)
	if putResp.PrevKv!=nil{
	   //打印该k前一次的v
	   fmt.Println("PrevValue:",string(putResp.PrevKv.Value))
	}
   }
}
//初始化etcd与获取kv
func initAndGetKV()  {
   var (
	config clientv3.Config
	client *clientv3.Client
	err error
	kv clientv3.KV
	getResp *clientv3.GetResponse
	v *mvccpb.KeyValue
   )
   //初始化 start-------
   //客户端配置
   config = clientv3.Config{
	Endpoints:[]string{"127.0.0.1:2379"},
	DialTimeout:5*time.Second,
   }
   //建立连接
   if client,err = clientv3.New(config);err!=nil{
	fmt.Println(err)
	return
   }
   //初始化 end-------
   //用于读写etcd的kv键值对
   kv = clientv3.NewKV(client)
   //获取指定k的值
   //如果想获取指定前缀的所有k，则参数为:context.TODO(),"/test/",clientv3.WithPrefix()
   if getResp,err = kv.Get(context.TODO(),"/test/job1");err!=nil{
	fmt.Println(err)
   }else{
	//查看
	//getResp.Kvs为[]*mvccpb.KeyValue,每个切片包含key、create_revision、mod_revision、value
	//getResp.Count为查询的是数量
	fmt.Println(getResp.Kvs,getResp.Count)
	//需要通过遍历来具体获取值
	for _,v = range getResp.Kvs{
	   fmt.Println(string(v.Key))
	   fmt.Println(string(v.Value))
	   fmt.Println(v.CreateRevision)
	   fmt.Println(v.ModRevision)
	}
	
   }
}
//初始化etcd与删除kv
func initAndDeleteKV()  {
   var (
	config clientv3.Config
	client *clientv3.Client
	err error
	kv clientv3.KV
	delResp *clientv3.DeleteResponse
	v *mvccpb.KeyValue
   )
   //初始化 start-------
   //客户端配置
   config = clientv3.Config{
	Endpoints:[]string{"127.0.0.1:2379"},
	DialTimeout:5*time.Second,
   }
   //建立连接
   if client,err = clientv3.New(config);err!=nil{
	fmt.Println(err)
	return
   }
   //初始化 end-------
   //用于读写etcd的kv键值对
   kv = clientv3.NewKV(client)
   //删除k（clientv3.WithPrevKV()为删除前获取旧的值）
   //如果想删除指定前缀的所有k，则参数为:context.TODO(),"/test/",clientv3.WithPrefix()
   //以后会添加说明:clientv3.WithFromKey(),clientv3.WithLimit(2)
   if delResp,err = kv.Delete(context.TODO(),"/test/job1",clientv3.WithPrevKV());err!=nil{
	fmt.Println(err)
   }else{
	//查看被删除之前是什么
	if len(delResp.PrevKvs) != 0 {
	   for _,v = range delResp.PrevKvs{
		fmt.Println(string(v.Key),string(v.Value))
	   }
	}
	
   }
   
}
//初始化etcd与租约挂钩
func initAndApplyLease ()  {
   var (
	config clientv3.Config
	client *clientv3.Client
	err error
	lease clientv3.Lease
	leaseGrantResp *clientv3.LeaseGrantResponse
	leaseId  clientv3.LeaseID
	kv clientv3.KV
	putResp *clientv3.PutResponse
   )
   //初始化 start-------
   //客户端配置
   config = clientv3.Config{
	Endpoints:[]string{"127.0.0.1:2379"},
	DialTimeout:5*time.Second,
   }
   //建立连接
   if client,err = clientv3.New(config);err!=nil{
	fmt.Println(err)
	return
   }
   //初始化 end-------
   //申请一个租约
   lease = clientv3.NewLease(client)
   //申请一个10s的租约
   if leaseGrantResp,err = lease.Grant(context.TODO(),10);err!=nil{
	fmt.Println(err)
	return
   }
   //拿到租约id
   leaseId = leaseGrantResp.ID
   
   //创建一个k，让其与租约相关联，从而实现10s后自动过期
   kv = clientv3.NewKV(client)
   if putResp,err = kv.Put(context.TODO(),"/test/job3","hhh",clientv3.WithLease(leaseId));err!=nil{
	fmt.Println(err)
	return
   }
   fmt.Println("写入成功,",putResp.Header.Revision)
   //定时看一下key过期了没有
	for {
		getResp,err := kv.Get(context.TODO(),"/test/job3")
		if err!=nil{
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("没有过期",getResp.Kvs)
		time.Sleep(2*time.Second)
	}
	fmt.Println("ok!")
   
}
//租约续租
func initAndRenewalLease ()  {
var (
	config clientv3.Config
	client *clientv3.Client
	err error
	lease clientv3.Lease
	leaseGrantResp *clientv3.LeaseGrantResponse
	leaseId  clientv3.LeaseID
	keepResp *clientv3.LeaseKeepAliveResponse
	keepRespChan <- chan *clientv3.LeaseKeepAliveResponse
	kv clientv3.KV
	putResp *clientv3.PutResponse
   )
   //初始化 start-------
   //客户端配置
   config = clientv3.Config{
	Endpoints:[]string{"127.0.0.1:2379"},
	DialTimeout:5*time.Second,
   }
   //建立连接
   if client,err = clientv3.New(config);err!=nil{
	fmt.Println(err)
	return
   }
   //初始化 end-------
   //申请一个租约
   lease = clientv3.NewLease(client)
   //申请一个10s的租约
   if leaseGrantResp,err = lease.Grant(context.TODO(),10);err!=nil{
	fmt.Println(err)
	return
   }
   //拿到租约id
   leaseId = leaseGrantResp.ID
   
   //自动续租
   ctx,_ := context.WithTimeout(context.TODO(),5*time.Second)
   //续租5s，就停止了续租，再加上10s的TTL，总共生存期为15s
   //5s后取消自动续租
   if keepRespChan ,err = lease.KeepAlive(ctx,leaseId);err!=nil{
	fmt.Println(err)
	return
   }
   //处理续约应答的协程
   go func() {
	for{
	   select {
		case keepResp = <-keepRespChan:
			if keepResp==nil{
				fmt.Println("租约已经失效")
				goto END
			}else {
				//每秒会续租一次，因此会收到一次应答
				fmt.Println("收到自动续租的应答",keepResp.ID)
			}
		 }
	     }
	     END:
	}()
   
   
   //创建一个k，让其与租约相关联，从而实现10s后自动过期
   kv = clientv3.NewKV(client)
   if putResp,err = kv.Put(context.TODO(),"/test/job3","hhh",clientv3.WithLease(leaseId));err!=nil{
	fmt.Println(err)
	return
   }
   fmt.Println("写入成功,",putResp.Header.Revision)
   //定时看一下key过期了没有
	for {
		getResp,err := kv.Get(context.TODO(),"/test/job3")
		if err!=nil{
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("没有过期",getResp.Kvs)
		time.Sleep(2*time.Second)
	}
	fmt.Println("ok!")
}
//监听kv变化
func initAndWatchKV(){
 var (
	config clientv3.Config
	client *clientv3.Client
	err error
	kv clientv3.KV
   )
   //初始化 start-------
   //客户端配置
   config = clientv3.Config{
	Endpoints:[]string{"127.0.0.1:2379"},
	DialTimeout:5*time.Second,
   }
   //建立连接
   if client,err = clientv3.New(config);err!=nil{
	fmt.Println(err)
	return
   }
   //初始化 end-------
   
   kv = clientv3.NewKV(client)
   
   //模拟etcd中KV的变化
   go func() {
	kv.Put(context.TODO(),"/test/job7","i am job7")
	kv.Delete(context.TODO(),"/test/job7")
	time.Sleep(1*time.Second)
   }()
   
   //创建5s的上下文
   ctx,cancelFunc := context.WithCancel(context.TODO())
   time.AfterFunc(5*time.Second, func() {
	cancelFunc()
   })
   //先Get到当前的值，再去监听后续的变化
   getResp,err := kv.Get(ctx,"/test/job7")
   if err!=nil{
	fmt.Println(err)
	return
   }
   //现在key是存在的
   if len(getResp.Kvs)!=0{
	fmt.Println("当前值:",string(getResp.Kvs[0].Value))
   }
   //当前etcd集群事务id,单调递增
   //从当前事务往后开始进行watch
   watchStartRevision := getResp.Header.Revision +1
   //创建watcher
   watcher := clientv3.NewWatcher(client)
   fmt.Println("从该版本往后开始监听:",watchStartRevision)
   //启动监听
   watchRespChan := watcher.Watch(context.TODO(), "/test/job7", clientv3.WithRev(watchStartRevision))
   //处理kv变化的事件
   for watchResp := range watchRespChan{
	for _,event := range watchResp.Events{
		switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为:",string(event.Kv.Value),"Revision:",event.Kv.CreateRevision,event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了","Revision:",event.Kv.ModRevision)
			}
	}
   }
	fmt.Println("ok!")
}
//使用OpAction来代替kv.Action
func initAndOp(){
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		putOp clientv3.Op
		opResp clientv3.OpResponse
	)
	//客户端配置
	config = clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:5*time.Second,
	}
	//建立连接
	if client,err = clientv3.New(config);err!=nil{
		fmt.Println(err)
		return
	}
	kv = clientv3.NewKV(client)
	//Op:是client内部对kv.Action的抽象
	putOp = clientv3.OpPut("/test/job8","i am job8")
	//执行Op
	if opResp,err=kv.Do(context.TODO(),putOp);err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("写入Revision：",opResp.Put().Header.Revision)
	//创建op
	getOp := clientv3.OpGet("/test/job8")
	//执行op
	if opResp,err = kv.Do(context.TODO(),getOp);err!=nil{
		fmt.Println(err)
		return
	}
	//打印
	fmt.Println("数据Revision：",opResp.Get().Kvs[0].ModRevision)//create rev == mod rev
	fmt.Println("数据Value：",string(opResp.Get().Kvs[0].Value))
}
//使用op操作和txn事务(if else then)来实现乐观锁
//lease实现锁的自动过期
func initAndTxn()  {
	var (
	   config clientv3.Config
	   client *clientv3.Client
	   err error
	   lease clientv3.Lease
	   leaseGrantResp *clientv3.LeaseGrantResponse
	   leaseId clientv3.LeaseID
	   keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
	   keepResp *clientv3.LeaseKeepAliveResponse
	   ctx context.Context
           cancelFunc context.CancelFunc
	   kv clientv3.KV
	   txn clientv3.Txn
	   txnResp *clientv3.TxnResponse
	)
	//客户端配置
	config = clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:5*time.Second,
	}
	//建立连接
	if client,err = clientv3.New(config);err!=nil{
	   fmt.Println(err)
	   return
	}
	//1.上锁(创建租约，自动续租，拿着租约去抢占一个key)
	lease = clientv3.NewLease(client)
	//申请一个5s的租约
	if leaseGrantResp,err = lease.Grant(context.TODO(),5);err!=nil{
	   fmt.Println(err)
	   return
	}
	//获取租约id
	leaseId = leaseGrantResp.ID
	//准备一个用于取消自动续租的context
	ctx,cancelFunc = context.WithCancel(context.TODO())
	//5s后会取消自动续租
	if keepRespChan,err=lease.KeepAlive(ctx,leaseId);err!=nil{
	   fmt.Println(err)
	   return
	}
	//处理续租应答的协程
	go func() {
	   for{
		select {
		   case keepResp = <-keepRespChan:
			if keepResp==nil{
			   fmt.Println("租约已经失效")
			   goto END
			}else {
			   //每秒会续租一次，因此会收到一次应答
			   fmt.Println("收到自动续租的应答",keepResp.ID)
			}
		  }
	    }
	    END:
	}()
	//if 不存在key，then设置它，else抢锁失败
	kv = clientv3.NewKV(client)
	//创建事务
	txn = kv.Txn(context.TODO())
	//定义事务
	txn.If(clientv3.Compare(
	    //如果key不存在
	    clientv3.CreateRevision("/test/job9"),"=",0)).
	    //创建key
	    Then(clientv3.OpPut("/test/job9","",clientv3.WithLease(leaseId))).
	    //否则抢锁失败
	    Else(clientv3.OpGet("/test/job9"))
	//提交事务
	if txnResp,err=txn.Commit();err!=nil{
	   fmt.Println(err)
	   return//没有问题，接下来会执行defer
	}
	//判断是否抢到了锁
	if !txnResp.Succeeded{
	    fmt.Println("锁被占用了:",string(
	        txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
	}
	//2.处理业务
	//在锁内，很安全
	fmt.Println("处理任务")
	time.Sleep(time.Second*5)
	//3.释放锁(取消自动续租)
	//defer会把租约释放掉，关联的kv就被删除了
	//确保函数退出后，自动续租会停止
	defer cancelFunc()
	defer lease.Revoke(context.TODO(),leaseId)
}
```

#### mongo入门操作
go get go.mongodb.org/mongo-driver/mongo
用法API可以参考文章[https://www.jianshu.com/p/0344a21e8040] 
```go
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//初始化mongo
func initConn()(*mongo.Collection){
   	var(
	     ctx context.Context
	     client *mongo.Client
	     err error
	     collection *mongo.Collection
	)
	//1.建立连接
	ctx,_ = context.WithTimeout(context.TODO(),time.Second*5)
	if client,err=
		mongo.Connect(ctx,options.Client().ApplyURI("mongodb://127.0.0.1:27017"));err!=nil{
		fmt.Println(err)
		return
	}
	//2.选择数据库和表collection
	collection = client.Database("cron").Collection("log")
	return collection
}
//任务的执行时间点
type TimePoint struct {
     StartTime int64 `bson:"start_time"`
     EndTime int64 `bson:"end_time"`
}
//一条日志
type LogRecord struct {
     JobName string `bson:"job_name"`//任务名
     Command string `bson:"command"`//shell命令
     Err string `bson:"err"`//脚本错误
     Content string `bson:"content"`//脚本输出
     TimePoint TimePoint `bson:"time_point"`//执行时间点
}
//插入记录
func insertOneToMongo(){
   	var(
	     err error
	     collection *mongo.Collection
	     record *LogRecord
	     result *mongo.InsertOneResult
	     docId primitive.ObjectID
	)
	collection = initConn()
	//插入记录(bson)
	record = &LogRecord{
		"job11",
		"echo hello world",
		"",
		"Hello world",
		TimePoint{
			StartTime:time.Now().Unix(),
			EndTime:time.Now().Unix()+10,
		},
	}
	if result,err = collection.InsertOne(context.TODO(),record);err!=nil{
		fmt.Println(err)
		return
	}
	//_id:默认生成一个全局唯一的ID值,ObjectID(12字节的二进制)
	docId = result.InsertedID.(primitive.ObjectID)
	fmt.Println("自增id：",docId.Hex())
}
//jobName过滤条件
type FindByJobName struct {
     JobName string `bson:"job_name"`
}
//查询记录
func findToMongo(){
	var(
	     err error
	     collection *mongo.Collection
	     record *LogRecord
	     cond *FindByJobName
	     cursor *mongo.Cursor
	)
	collection = initConn()
	
	//按照job_name字段过滤，想找出job_name为10，找出5条
	cond = &FindByJobName{JobName:"job10"}//{"job_name":"job10"}
	//查询(过滤+翻页参数)
	if cursor,err = collection.Find(
		context.TODO(),
		cond,//过滤条件
		options.Find().SetSkip(0),//设置跳过第几条
		options.Find().SetLimit(2));//设置获取几条
	   err!=nil{
		fmt.Println(err)
		return
	}
	//释放游标
	defer cursor.Close(context.TODO())
	//遍历结果集
	for cursor.Next(context.TODO()){
	    //定义一个日志对象
	    record = &LogRecord{}
	    //反序列化bson到对象
	    if err = cursor.Decode(record);err!=nil{
		fmt.Println(err)
		return
	     }
	     //把日志打印出来
	     fmt.Println(*record)
	}
	
}
//删除数据
//start_time小于某时间
//{"$lt":timestamp
