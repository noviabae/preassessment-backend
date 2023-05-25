SELECT t1.id, t1.name, t2.name AS parent_name
FROM my_table t1
LEFT JOIN my_table t2 ON t1.parent_id = t2.id