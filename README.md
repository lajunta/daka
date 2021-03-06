## 打卡分析程序

本程序用来对疫情期间学生的打卡数据进行分析

- 统计学生打卡的日期
- 统计学生未打卡的日期
- 统计学生共打卡多少次
- 检查学生是否按指定日期连续打卡

### 使用说明

在程序运行的目录下至少需要三个文件

1. current.csv 全校所有学生全部打卡的实际数据,前三个字段为打卡时间,学生班级,学生姓名, 顺序不能错
1. student.csv 全校学生的花名册,需要三个字段,班级,姓名,学号,并且同一个班级不要有重名
1. dates.txt 规定打卡的日期文件,格式可为 2020/2/20, 一行为须打卡的日期, 一行为写一天

### 运行

确保以上三个数据文件在程序运行的目录下面

点击 `打卡程序` ,会在当前目录下生成output.csv 文件

会生成学生的打卡统计数据

字段分别为：

```
班级,姓名,学号,实际打卡次数,打卡日期,缺少日期,缺少天数,连续打卡
```

以上数据可做为对班级的分析,也可做为开学的数据支持