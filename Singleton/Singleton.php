<?php
// 单例模式
class Singleton {
    // 创建静态私有的变量保存该类对象
    private static $instance;

    // 防止使用new直接创建对象
    private function __construct(){}

    // 防止使用clone克隆对象
    private function __clone(){}

    // 获取该类对象
    public static function getInstance() {
        // 判断$instance是否是Singleton的对象, 如果不是则创建; 如果是, 则直接返回
        if (! self::$instance instanceof self) {
            self::$instance = new Singleton();
        }

        return self::$instance;
    }
    public function say() {
        echo "Hello World!";
    }
}

// 实例
$s = Singleton::getInstance ();
$s->say ();

// $s2 = new Singleton(); // error
// $s3 = clone $s1; // error