# Design

## download

负责下载 url 对应的网页

## process

负责处理下载后的网页内容
1. 处理 json 返回值，解析节点中的数组

## pipeline

负责将处理后的内容输出到各种源(控制台、文件、消息队列、数据库等)

## schedule

任务调度

## register

注册爬虫

## http

REST 接口

## spider

整合爬虫各模块

## core

核心，负责生命周期，上下文管理