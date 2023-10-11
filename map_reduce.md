## Настройка MapReduce

Копируем бинари:
```bash
cd map/
go build
cd ../reduce/
go build
cd ..
scp ./map/map team3@91.185.86.24:/tmp/
scp ./reduce/reduce team3@91.185.86.24:/tmp/
```

```bash
vim /usr/local/hadoop/etc/hadoop/mapred-site.xml

...
<!-- Put site-specific property overrides in this file. -->

<configuration>
   <property>
      <name>mapreduce.framework.name</name>
      <value>yarn</value>
   </property>
</configuration>
```

```bash
vim /usr/local/hadoop/etc/hadoop/yarn-site.xml

...
<configuration>

<!-- Site specific YARN configuration properties -->
  <property>
    <name>yarn.nodemanager.aux-services</name>
    <value>mapreduce_shuffle</value>
  </property>
</configuration>
```

```bash
hadoop jar $HADOOP_HOME/share/hadoop/tools/lib/hadoop-streaming-*.jar  -files /tmp/map,  /tmp/reduce  -input input/ncdc/all  -output output  -mapper /tmp/map -combiner /tmp/reduce  -reducer /tmp/reduce
```
