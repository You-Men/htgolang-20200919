1. 初始化命令
    flag -init
    a. 已初始化，提示是否重新初始化
    b. 初始管理信息
        添加一个用户
    c. 初始化存储方式: gob/csv/json


2. 用户管理
    gob
    csv
    json

    滚动存储: 保存最后5g文件(保存最后5次操作的文件)

    a. user.csv
    b. user.csv.1
    c. user.csv.2
    d. user.csv.4
    e. user.csv.5
    f. user.csv.6

    固定文件
    save():
        检查文件是否存在
            赋值 另外一个文件名 => filename.unixtime
        filename.unixtime 格式的文件数量 4
        文件数量 > 4:
            排序
            移除文件

    动态：
    save():
        保存文件 => filename.unixtime
        filename.unixtime 格式的文件数量 5
        文件数量 > 5:
            排序
            移除文件

    load():

3. tailf => 收集日志
    tail -f
