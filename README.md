# IntroductiontoAlgorithms
算法导论

# 一、算法
算法的定义：任何良定义(well-define:定义明确、无歧义)的计算过程。

算法的正确性：对每个符合条件的输入实例，都以正确的输出停机

算法问题的共性：
1. 存在很多候选解
2. 存在实际应用

数据结构： 存储和组织数据的方式，旨在便于访问和修改

算法效率： 一般量度是速度，即，一个算法花费多长时间产生结果


循环不变式(插入排序有示例)：

性质：
1. 初始化：循环的第一次迭代之前，为true
2. 保持：如果循环的某次迭代之前它为true，那么下次迭代它仍为true 
3. 终止：循环终止时，不变式为我们提供一个有用的性质。并且该性质有助于证明算法的正确性

# 二、分析算法

含义：预测算法需要的资源，除了内存，通信外，通常想度量的是时间

算法需要的时间与输入规模同步增长，所以通常把一个程序的运行时间描述成其输入规模的函数

运行时间：执行的基本操作数或步数

往往只求最坏情况运行时间：
1. 最坏情况运行时间给出了任何输入的运行时间的一个上界
2. 某些算法，最坏情况经常出现
3. “平均情况” 往往与最坏情况大致一样差

# 三、 分治法
增量法：
插入排序

分治法：归并排序

分治法思想：将原问题分解为几个规模较小但类似于原问题的子问题

步骤：
1. 分解：原问题分解为若干子问题
2. 解决：递归的求解各个子问题
3. 合并：子问题的解成原问题的解


# 四、函数的增长


# Tips
1. 整数左移一位等价于该整数乘以2
