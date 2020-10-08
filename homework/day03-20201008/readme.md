1. 用户管理功能

// 用户增删改查
定义全局变量
users := []map[string]string{}

每个元素
    ID
    名称
    联系方式
    通信地址

4个go程序,每个程序写一个功能
1. 增加
    add函数
    从命令行分别输入名称、联系方式、通信地址

    生成ID => 查找users中最大的id+1（无元素id=1）
    放入到users

    fmt.Println(users)

2. 删除
    del函数
    从命令行中输入要删除的用户ID
    验证ID是否存在，如果存在，打印需要删除用户信息
    并让用户输入y/n确认是否删除
    输入y删除用户信息

    fmt.Println(users)

3. 修改
    modify函数
    从命令行输入要修改的用户ID
    验证ID是否存在，如果存在，打印需要删除用户信息
    并让用户输入y/n确认是否修改
    输入y修改用户信息，继续让从命令行分别输入
    用户名，联系方式，地址 进行更新

    fmt.Println(users)

4. 查找
    query函数
    从命令行输入要查询的字符串
    遍历比较用户的名称，地址，联系方式，包含要查找的字符串就进行输出


    Id：xxx
    名字：xxx