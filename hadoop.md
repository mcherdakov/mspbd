## Установка hadoop

```bash
sudo apt update && sudo apt upgrade
```

```bash
sudo iptables -I INPUT -p tcp --dport 9870 -j ACCEPT
sudo iptables -I INPUT -p tcp --dport 8020 -j ACCEPT
sudo iptables -I INPUT -p tcp --match multiport --dports 9866,9864,9867 -j ACCEPT
```

```bash
sudo apt install iptables-persistent
sudo netfilter-persistent save
```


```bash
sudo vim /etc/hosts
10.0.10.5 main
10.0.10.11 worker1
10.0.10.12 worker2
```

```bash
sudo apt install default-jdk
```


```bash
wget https://dlcdn.apache.org/hadoop/common/hadoop-3.3.1/hadoop-3.3.1.tar.gz

sudo mkdir /usr/local/hadoop && sudo tar -zxf hadoop-*.tar.gz -C /usr/local/hadoop --strip-components 1

sudo useradd hadoop -m

sudo passwd hadoop -> hadoop

sudo chown -R hadoop:hadoop /usr/local/hadoop
```

```bash
sudo vim /etc/profile.d/hadoop.sh
export HADOOP_HOME=/usr/local/hadoop
export HADOOP_HDFS_HOME=$HADOOP_HOME
export HADOOP_MAPRED_HOME=$HADOOP_HOME
export HADOOP_COMMON_HOME=$HADOOP_HOME
export HADOOP_COMMON_LIB_NATIVE_DIR=$HADOOP_HOME/lib/native
export HADOOP_OPTS="$HADOOP_OPTS -Djava.library.path=$HADOOP_HOME/lib/native"
export YARN_HOME=$HADOOP_HOME
export PATH="$PATH:${HADOOP_HOME}/bin:${HADOOP_HOME}/sbin"
```

```bash
sudo vim /usr/local/hadoop/etc/hadoop/hadoop-env.sh
export JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64
```

```bash
sudo vim /usr/local/hadoop/etc/hadoop/core-site.xml
<configuration>
   <property>
      <name>fs.default.name</name>
      <value>hdfs://localhost:9000</value>
   </property>
</configuration>
```

```bash
sudo vim /usr/local/hadoop/etc/hadoop/hdfs-site.xml
<configuration>
   <property>
      <name>dfs.replication</name>
      <value>1</value>
   </property>
   <property>
      <name>dfs.name.dir</name>
      <value>file:///hadoop/hdfs/namenode</value>
   </property>
   <property>
      <name>dfs.data.dir</name>
      <value>file:///hadoop/hdfs/datanode</value>
   </property>
</configuration>
```

```bash
sudo mkdir -p /hadoop/hdfs/{namenode,datanode}
sudo chown -R hadoop:hadoop /hadoop
```

```bash
ssh-keygen
cat .ssh/id_rsa.pub > .ssh/authorized_keys
```

```bash
/usr/local/hadoop/bin/hdfs namenode -format
/usr/local/hadoop/sbin/start-dfs.sh
/usr/local/hadoop/sbin/start-yarn.sh
```

## Создание файла
```bash
hdfs dfs -mkdir /tmp
echo 'Hello world!' > test.txt
hdfs dfs -put test.txt /tmp
```

Смотрим размер через утилиту hdfs
```bash
hdfs dfs -du -h /tmp
14  14  /tmp/test.txt
```

Смотрим размер в хостовой файловой системе
```bash
sudo du -ah .
4.0K  ./hdfs/namenode/in_use.lock
4.0K  ./hdfs/namenode/current/fsimage_0000000000000000000
4.0K  ./hdfs/namenode/current/edits_0000000000000000001-0000000000000000002
4.0K  ./hdfs/namenode/current/fsimage_0000000000000000000.md5
4.0K  ./hdfs/namenode/current/VERSION
1.0M  ./hdfs/namenode/current/edits_inprogress_0000000000000000003
4.0K  ./hdfs/namenode/current/seen_txid
1.1M  ./hdfs/namenode/current
1.1M  ./hdfs/namenode
4.0K  ./hdfs/datanode/in_use.lock
4.0K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674/tmp
4.0K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674/scanner.cursor
4.0K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674/current/finalized/subdir0/subdir0/blk_1073741825
4.0K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674/current/finalized/subdir0/subdir0/blk_1073741825_1001.meta
12K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674/current/finalized/subdir0/subdir0
16K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674/current/finalized/subdir0
20K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674/current/finalized
4.0K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674/current/rbw
4.0K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674/current/VERSION
32K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674/current
44K  ./hdfs/datanode/current/BP-1093842354-127.0.1.1-1696440454674
4.0K  ./hdfs/datanode/current/VERSION
52K  ./hdfs/datanode/current
60K  ./hdfs/datanode
1.1M  ./hdfs
1.1M  .
```
















