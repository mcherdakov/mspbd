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


