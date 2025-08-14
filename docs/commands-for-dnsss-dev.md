# commands-for-dnsss-dev

这里是一些在开发 dnsss 项目时需要用到的命令。



## Linux 查看端口占用情况

    sudo lsof -i:[端口号]


## 查看 systemd 中的 dns 服务

    systemctl status systemd-resolved


## 关闭 systemd 中的 服务
 
    systemctl stop systemd-resolved
 

## 允许监听linux  端口

临时生效：

    sysctl net.ipv4.ip_unprivileged_port_start=0

永久生效：

    echo "net.ipv4.ip_unprivileged_port_start=0" >> /etc/sysctl.conf
    sysctl -p
