# gApi

---

### **1. [Install Go](https://go.dev/doc/install)**
### **2.** Install the dependencies
```bash go run tidy```

### **3.** Run the migrations once the models are in place
```bash go mod migrate/migrate.go ```

### **4.** Run the main Go file once the DB is setup (below).
```bash go run main ```


### **5.** curl localhost:8080

You can [View the endpoints here](ENDPOINTS.md)

---

## Setup Postgress

### **1. Update the Package List**
```bash
sudo apt update
sudo apt upgrade -y
```
This ensures your system has the latest information on available packages.

---

### **2. Install PostgreSQL and PostgreSQL Client**
Imho the best way to install **PostgreSQL** on Debian for local 
development is to just use the **Debian package manager (APT)**. 
This ensures you get stable, well-maintained versions that 
are compatible with the system. 

The default version in Debian's repository 
may be slightly behind the latest release, but for local development, it's usually sufficient.

```bash
sudo apt install -y postgresql postgresql-client
```
This installs the **PostgreSQL server** and **psql** (the command-line client). 

---

### **3. Verify PostgreSQL is Running**
Check the status of the PostgreSQL service. 

**NOTE:** The Debian way of calling this is:

```bash
 systemctl status postgresql@15-main
 ```


```bash
sudo systemctl status postgresql
```
Look for `active (running)` in the output. If it's not running, you can start it with:
```bash
sudo systemctl start postgresql
```

---

### **4. Set the PostgreSQL Service to Start on Boot**
```bash
sudo systemctl enable postgresql
```

---

### **5. Switch to the Default `postgres` User**
PostgreSQL creates a default system user called **`postgres`**. Switch to it:
```bash
sudo -i -u postgres
```

```

---

### **6. Create a New User and Database**
If you want to create a new database and user for your project:
```sql
CREATE USER myuser WITH PASSWORD 'mypassword';
CREATE DATABASE mydb OWNER myuser;
GRANT ALL PRIVILEGES ON DATABASE mydb TO myuser;
```

Exit the **psql** prompt with:
```sql
\q
```


Restart PostgreSQL to apply changes:
```bash
sudo systemctl restart postgresql
```

---

### **8. Connect to PostgreSQL**
You can now connect using your new user:
```bash
psql -U myuser -d mydb -h 127.0.0.1 -W
```
The `-W` flag tells `psql` to prompt you for the password.

---

### **Bonus: Useful Commands**
- **List all databases**: 
  ```sql
  \l
  ```

- **Switch to a specific database**:
  ```sql
  \c mydb
  ```

- **List all users**: 
  ```sql
  \du
  ```

- **List all tables in the current database**: 
  ```sql
  \dt
  ```

---

```bash
sudo systemctl status postgresql@15-main
sudo systemctl restart postgresql
sudo systemctl status postgresql@15-main
```


sudo -i -u postgres
psql


CREATE ROLE
postgres=# CREATE DATABASE mydb OWNER myuser;
ERROR:  role "myuser" does not exist
postgres=# CREATE DATABASE gApi OWNER tk;

\q



sudo systemctl list-units | grep postgresql


sudo journalctl -u postgresql


