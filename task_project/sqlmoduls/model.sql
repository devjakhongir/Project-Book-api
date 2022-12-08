create table books (
    id UUID default gen_random_uuid(),
    name varchar(24) not null,
    price int,
    author_name varchar(36) not null,
    publisher_date varchar,
    page int not null,
    janr varchar
    
    
);


insert into books(name,price, author_name,publisher_date,page,janr) values

('Otkan kunlar',36000,'Abdulla Qodiriy','01-01-1922',400,'roman'),
('Dunyoning ishlari',40000,'Otkir Hoshimov','02-07-2005',336,'qissa'),
('Ikki eshik orasi',50000,'Otkir Hoshimov','05-07-2012',624,'roman'),
('Shum bola',56000,'Gafur Gulom','01-01-1936',550,'qissa'),
('Ufq',34000,'Said Ahmad','12-12-1973',400,'roman');


--created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
--updated_at timestamp