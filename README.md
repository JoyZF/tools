# tools

toolkit by cobra, A collection of common tools,only for mac.

## feature
- [x] 快速打开应用
- [x] 获取时间戳
- [x] 获取日期
- [x] base64加解密
- [x] 查询天气预报
- [x] 获取当前IP
- [x] KFC文案
- [x] 获取今日老黄历（原值返回）
- [x] git push 自动化


## How to install
```shell
go install github.com/JoyZF/tools
```

## How to use

快速打开应用
```shell   
tools open -n c
```

获取时间戳
```shell
tools time
```

获取指定日期的时间戳
```shell
tools time "2024-01-01 00:00:00"
```

获取日期
```shell
tools date
``` 

时间戳转日期
```shell
tools date 1630483200
``` 

base64 加解密
```shell
tools base64 -d encode -s "需要加密的数据"
```

查询当前城市天气预报（根据IP）
```shell
tools weather
```

查询指定城市天气预报
```shell
tools weather 北京
```

获取当前IP
```shell
tools ip
```

随机获取KFC文案
```shell
tools kfc
```


随机今日老黄历
```shell
tools luck
```
 