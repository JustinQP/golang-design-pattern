[TOC]

## 一、原型模式

原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象。

原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置。

### (1) 模式动机

原型模式的动机在于创建对象的过程通常比较耗时或复杂，而且新创建的对象与已有对象之间往往只有细微的差别。此时，通过克隆已有对象来创建新对象，可以大大提高性能和降低复杂性。

### (2) 模式定义

原型模式是一种创建型设计模式，它通过克隆已有对象来创建新对象，避免了直接调用构造函数进行对象的创建。

#### (2.1) 模式结构

原型模式的主要结构包含以下几个角色：

- 原型（Prototype）：声明一个克隆自身的方法，克隆方法可以返回一个克隆后的对象。
- 具体原型（Concrete Prototype）：实现原型的克隆方法。
- 客户端（Client）：通过调用克隆方法来克隆原型对象，创建新的对象。

### (3) 模式分析

#### (3.1) 优点

- 通过克隆已有对象创建新对象，避免了直接调用构造函数，提高了对象的创建性能。
- 可以动态地添加或删除对象的部分，通过对已有对象进行克隆来创建新对象，并修改新对象的部分属性。
- 可以隐藏对象创建的细节，使得创建过程对客户端透明，简化了客户端的操作。

#### (3.2) 缺点

- 克隆对象需要实现Clone方法，这对于每个可克隆的类都是必须的。
- 如果对象包含引用类型的成员变量，需要考虑深克隆和浅克隆的问题。
- 克隆对象可能会破坏对象的封装性，因为克隆对象可以直接访问原型对象的私有成员变量。

#### (3.3) 适用场景

- 创建对象的过程比较耗时或复杂。
- 待创建的对象与已有对象之间只有细微的差别，克隆已有对象来创建新对象时，只需要修改差异部分即可。
- 需要动态地添加或删除对象的部分，通过克隆来创建新对象并修改差异部分。

### (4) 应用场景

原型模式在以下场景中常被使用：

- 创建对象的过程较为复杂，包含初始化和配置等，并且这些过程可以通过克隆已有对象来避免。
- 当需要创建一个对象的多个副本且每个副本都可能需要进行一些定制化的修改时。

### (5) 总结

原型模式通过克隆已有对象来创建新对象，避免了直接调用构造函数，提高了对象的创建性能和降低复杂性。它适用于创建对象过程复杂、对象之间差异较小以及需要动态添加或删除对象部分的场景。然而，使用原型模式需要注意克隆过程中对引用类型成员变量的处理，以及可能破坏对象封装性的问题。

