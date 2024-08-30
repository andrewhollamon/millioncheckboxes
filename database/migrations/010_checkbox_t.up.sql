/*
 CHECKBOX_T
(ranged partitioned on CHECKBOX_NBR:  https://www.postgresql.org/docs/current/ddl-partitioning.html)
- CHECKBOX_NBR int, unique, not null, primary key
- CHECKED_STATE boolean, not null, default false
- INDEXES
  - PK: CHECKBOX_NBR
 */

CREATE TABLE CHECKBOX_T (
    CHECKBOX_NBR      INT UNIQUE NOT NULL PRIMARY KEY,
    CHECKED_STATE     BOOLEAN NOT NULL DEFAULT FALSE
) PARTITION BY RANGE(CHECKBOX_NBR)
;
CREATE TABLE CHECKBOX_PART00_T PARTITION OF CHECKBOX_T
    FOR VALUES FROM (0) TO (100000)
;
CREATE TABLE CHECKBOX_PART01_T PARTITION OF CHECKBOX_T
    FOR VALUES FROM (100000) TO (200000)
;
CREATE TABLE CHECKBOX_PART02_T PARTITION OF CHECKBOX_T
    FOR VALUES FROM (200000) TO (300000)
;
CREATE TABLE CHECKBOX_PART03_T PARTITION OF CHECKBOX_T
    FOR VALUES FROM (300000) TO (400000)
;
CREATE TABLE CHECKBOX_PART04_T PARTITION OF CHECKBOX_T
    FOR VALUES FROM (400000) TO (500000)
;
CREATE TABLE CHECKBOX_PART05_T PARTITION OF CHECKBOX_T
    FOR VALUES FROM (500000) TO (600000)
;
CREATE TABLE CHECKBOX_PART06_T PARTITION OF CHECKBOX_T
    FOR VALUES FROM (600000) TO (700000)
;
CREATE TABLE CHECKBOX_PART07_T PARTITION OF CHECKBOX_T
    FOR VALUES FROM (700000) TO (800000)
;
CREATE TABLE CHECKBOX_PART08_T PARTITION OF CHECKBOX_T
    FOR VALUES FROM (800000) TO (900000)
;
CREATE TABLE CHECKBOX_PART09_T PARTITION OF CHECKBOX_T
    FOR VALUES FROM (900000) TO (1000000)
;
