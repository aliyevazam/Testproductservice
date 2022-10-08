CREATE TABLE IF NOT EXISTS  products(id serial PRIMARY key,name VARCHAR(30)not NULL,model TEXT,typeId int REFERENCES types(id),categoryId int REFERENCES categories(id),price FLOAT,amount INT);
