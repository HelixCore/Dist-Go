create table pagina ( 
    id_pagina INT primary key, 
    titulo VARCHAR(50), 
    descripcion VARCHAR(50), 
    url VARCHAR(1000) 
);

create table busqueda ( 
    id_busqueda INT primary key, 
    palabra VARCHAR(50) 
);

create table resalta(
    id_busqueda INT,
    id_pagina INT,
    primary key (id_busqueda,id_pagina),
    FOREIGN key (id_busqueda) REFERENCES busqueda(id_busqueda) on update cascade,
    FOREIGN key (id_pagina) REFERENCES pagina(id_pagina) on update cascade

);