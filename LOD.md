迪米特法则(Law of Demeter, LOD)，有时候也叫做最少知识原则(Least Knowledge Principle, LKP)，它的定义是: 一个软件实体应当尽可能少地与其它实体发生相互作用。迪米特法则的初衷在于降低类之间的耦合。

举个例子，拿教师点名来讲，体育老师需要清点班上学生人数，教师一般不是自己亲自去数，而是委托组长或班长等人去清点，即教师通过下达命令至班长要求清点人数:

![WeChat7b40e17b1360547fd74bbe35d967c42c.png](http://ww1.sinaimg.cn/large/9e58a4edly1gfoc125sxgj20vo0d6409.jpg)


如果去掉 GroupLeader 这个中间人角色，教师就会直接去清点人数，这样做会违反迪米特法则。

迪米特法则总结：

类定义时尽量内敛，少用 public 权限修饰符，尽量使用 private、protected。