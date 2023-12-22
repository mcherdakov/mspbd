## Airflow 

### Установка и запуск Airflow

Заводим пользователя:
```bash
sudo adduser airflow
sudo su airflow
```

Ставим Airflow::
```bash
export AIRFLOW_HOME=~/airflow
AIRFLOW_VERSION=2.8.0
PYTHON_VERSION="$(python --version | cut -d " " -f 2 | cut -d "." -f 1-2)"
CONSTRAINT_URL="https://raw.githubusercontent.com/apache/airflow/constraints-${AIRFLOW_VERSION}/constraints-${PYTHON_VERSION}.txt"
pip install "apache-airflow==${AIRFLOW_VERSION}" --constraint "${CONSTRAINT_URL}"
echo 'export PATH="${PATH}:/home/airflow/.local/bin"' > ~/.bashrc
```

Запускаем Airflow:
```bash
cd /home/airflow/airflow
airflow standalone
```
![telegram-cloud-photo-size-2-5201748447282191357-y](https://github.com/mcherdakov/mspbd/assets/96630344/1ef17c2e-3d90-4ea7-b09e-3814cd76ef71)
![telegram-cloud-photo-size-2-5201748447282191363-y](https://github.com/mcherdakov/mspbd/assets/96630344/e7e1ce91-8263-4f2f-a0c9-bfa4a7765e93)




### Загрузка DAG

Настройка на хосте с Airflow:
```bash
airflow@airflow:~/airflow$ cat airflow.cfg  | grep dags_folder
dags_folder = /home/airflow/airflow/dags
airflow@airflow:~/airflow$ mkdir dags
airflow@airflow:~/airflow$ pip install apache-airflow-providers-ssh
```

Перенос DAG'а с локальной машины на хост с Airflow:
```bash
scp -i /home/mb/.ssh/yc-cloud-id-b1gtqeqluh845u5l39v6-bagishovmikail /tmp/dag.py bagishovmikail@51.250.30.154:/tmp/
```

Как выглядит в UI:
<img width="1662" alt="image" src="https://github.com/mcherdakov/mspbd/assets/96630344/0efc9d00-7e72-4f76-a4e8-20796e5792ea">


### Запуск DAG'а

Жмем на скрине выше на кнопочку Trigger DAG и получаем ошибку:
```bash
[2023-12-22, 20:21:30 UTC] {taskinstance.py:2699} ERROR - Task failed with exception
Traceback (most recent call last):
  File "/home/airflow/.local/lib/python3.10/site-packages/airflow/models/taskinstance.py", line 433, in _execute_task
    result = execute_callable(context=context, **execute_callable_kwargs)
  File "/home/airflow/.local/lib/python3.10/site-packages/airflow/providers/ssh/operators/ssh.py", line 177, in execute
    with self.get_ssh_client() as ssh_client:
  File "/home/airflow/.local/lib/python3.10/site-packages/airflow/providers/ssh/operators/ssh.py", line 142, in get_ssh_client
    return self.hook.get_conn()
  File "/home/airflow/.local/lib/python3.10/site-packages/airflow/providers/ssh/hooks/ssh.py", line 346, in get_conn
    for attempt in Retrying(
  File "/home/airflow/.local/lib/python3.10/site-packages/tenacity/__init__.py", line 347, in __iter__
    do = self.iter(retry_state=retry_state)
  File "/home/airflow/.local/lib/python3.10/site-packages/tenacity/__init__.py", line 325, in iter
    raise retry_exc.reraise()
  File "/home/airflow/.local/lib/python3.10/site-packages/tenacity/__init__.py", line 158, in reraise
    raise self.last_attempt.result()
  File "/usr/lib/python3.10/concurrent/futures/_base.py", line 451, in result
    return self.__get_result()
  File "/usr/lib/python3.10/concurrent/futures/_base.py", line 403, in __get_result
    raise self._exception
  File "/home/airflow/.local/lib/python3.10/site-packages/airflow/providers/ssh/hooks/ssh.py", line 353, in get_conn
    client.connect(**connect_kwargs)
  File "/home/airflow/.local/lib/python3.10/site-packages/paramiko/client.py", line 485, in connect
    self._auth(
  File "/home/airflow/.local/lib/python3.10/site-packages/paramiko/client.py", line 819, in _auth
    raise SSHException("No authentication methods available")
paramiko.ssh_exception.SSHException: No authentication methods available
```


### Починка DAG'а

Хотим починить поход по ssh. Для начала мы перезапустили Airflow, чтобы Admin->Connections пророс тип Summer Flow - SSH.

Теперь научим Airflow ходить на localhost по ssh:
```bash
ssh-keygen
cp ~/.ssh/id_rsa.pub ~/.ssh/authorized_keys
```

Варим нвоый Connection:
<img width="1645" alt="image" src="https://github.com/mcherdakov/mspbd/assets/96630344/b388edf7-610a-46b3-9e55-90b6dd44ace1">

Перезапускаем DAG и получаем новую ошибку:
```bash
[2023-12-22, 20:40:38 UTC] {ssh.py:478} INFO - Running command: 
curl -o out_file.zip 'URL'
[2023-12-22, 20:40:38 UTC] {ssh.py:529} WARNING -   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
[2023-12-22, 20:40:38 UTC] {ssh.py:529} WARNING -                                  Dload  Upload   Total   Spent    Left  Speed
[2023-12-22, 20:40:38 UTC] {ssh.py:529} WARNING - 
[2023-12-22, 20:40:38 UTC] {ssh.py:529} WARNING -   0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
[2023-12-22, 20:40:38 UTC] {ssh.py:529} WARNING - curl: (6) Could no
[2023-12-22, 20:40:38 UTC] {ssh.py:529} WARNING - t resolve host: URL
[2023-12-22, 20:40:38 UTC] {taskinstance.py:2699} ERROR - Task failed with exception
Traceback (most recent call last):
  File "/home/airflow/.local/lib/python3.10/site-packages/airflow/models/taskinstance.py", line 433, in _execute_task
    result = execute_callable(context=context, **execute_callable_kwargs)
  File "/home/airflow/.local/lib/python3.10/site-packages/airflow/providers/ssh/operators/ssh.py", line 178, in execute
    result = self.run_ssh_client_command(ssh_client, self.command, context=context)
  File "/home/airflow/.local/lib/python3.10/site-packages/airflow/providers/ssh/operators/ssh.py", line 166, in run_ssh_client_command
    self.raise_for_status(exit_status, agg_stderr, context=context)
  File "/home/airflow/.local/lib/python3.10/site-packages/airflow/providers/ssh/operators/ssh.py", line 160, in raise_for_status
    raise AirflowException(f"SSH operator error: exit status = {exit_status}")
airflow.exceptions.AirflowException: SSH operator error: exit status = 6
```
