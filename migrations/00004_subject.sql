-- +goose Up
CREATE TABLE subjects (
      id SERIAL PRIMARY KEY,
      name VARCHAR(100) NOT NULL,
      description TEXT
);

CREATE TABLE student_subject (
     student_id INT REFERENCES students(id),
     subject_id INT REFERENCES subjects(id),
     PRIMARY KEY (student_id, subject_id)
);

-- +goose Down
DROP TABLE IF EXISTS student_subject;
DROP TABLE IF EXISTS subjects;
