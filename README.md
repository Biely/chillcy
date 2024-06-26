# chillcy-rate-converter 汇率转换功能方法

> 实现一个输入币种与价格，输出人民币对应价格的方法

## 1、设计思路

### 1.1 背景

该功能可能应用于一个跨境电商服装平台，平台商品价格信息应该是通过抓取第三方平台而获得，需要保证币种判断的准确性和价格输出的准确性，并将此价格按人民币汇率转换后提供给国内消费者，根据此信息归纳了以下几个功能要点：

- 汇率是动态变化的，考虑保证每次转换都是通过接口请求最新汇率

- 当前业务仅需要五个币种转换，需要判断输入的币种是否符合当前功能支持范围

- 随着业务扩展，支持转换的货币种类可能增加，需要方便支持扩展更多货币换算

- 转换的目标货币也可能不仅限于人民币，也可能在转换的过程对价格进一步处理（折扣、运费...），能够灵活配置价格转换的算法

- 由于汇率接口得到的数据是基于美元的汇率，所以还需要进一步转换成对应人民币的汇率

### 1.2 实现方案

- 设计方案采用策略模式

- 用 golang 的 viper 库实现汇率接口地址、支持转换货币的灵活配置，创建一个 `config.yaml`的文件，程序执行时通过 viper 库读取配置文件，将接口地址、系统支持换算币种赋值给对应的全局变量

- 将正则提取币种和价格的函数 `MatchCurrencyAndPrice()`、请求接口并返回输入货币兑人民币汇率的函数 `GetRateToCNY()`封装进工具包 `utils`

## 2、开发环境

```
- golang版本 >= v1.19
- IDE：vsCode
- 操作系统： win10
```

## 3、运行项目

```
# 克隆项目
git clone https://github.com/Biely/chillcy.git

# 进入chillcy文件夹
cd chillcy

# 使用 go mod 并安装go依赖包
go generate

# 执行
go run main.go

# 编译
go build -o server main.go (windows编译命令为go build -o server.exe main.go )

# 运行二进制
./server (windows运行命令为 server.exe)
```

## 4、目录结构

```
    ├── config          (配置包)
    ├── global          (全局对象)                    
    ├── initialize      (初始化)                                                 
    ├── reteConverter   (汇率转换包)                                         
    ├── utils           (工具包)
    └── config.yaml     (配置文件)
```  