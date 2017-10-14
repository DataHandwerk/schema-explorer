-- mssql example db for regression tests
-- schema must match test code's expectations

if object_id('DataTypeTest', 'U') is not null
begin
	drop table DataTypeTest;
end
-- split here
create table DataTypeTest (
	intpk integer primary key,
	colCount integer,
	field_INT int,
	field_varcharmax varchar(max),
	field_nvarchar nvarchar(123),
	field_uniqueidentifier UNIQUEIDENTIFIER
);

delete DataTypeTest;
insert into DataTypeTest (
	intpk,
	colCount,
	field_INT,
	field_varcharmax,
	field_nvarchar,
	field_uniqueidentifier
) values (
	10, --intpk
	6, --colCount
	20, --field_INT
	'this is a ''text'' field',
	'blue',
	'b7a16c7a-a718-4ed8-97cb-20ccbadcc339'
),(
	11, --intpk
	0, --colCount
	-33, --field_INT
	'this is a ''text'' field',
	'blue',
	'b470fa05-2111-46f9-9c97-f103b594c5f0'
)
;
select * from DataTypeTest;
