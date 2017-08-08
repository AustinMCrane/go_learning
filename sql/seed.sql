-- Add some generic data
INSERT INTO guest (first_name, last_name)
VALUES
  ('Austin', 'Crane'),
  ('Jerek', 'Shoemaker'),
  ('Lakisha', 'Hammond');

INSERT INTO location (name)
VALUES
  ('Ceremony'),
  ('Reception');

INSERT INTO post (body, image)
VALUES
  ('This is the first post', 'https://scontent.fmci1-1.fna.fbcdn.net/v/t31.0-8/19143342_10154823871378458_1487953269318551064_o.jpg?oh=5b20bd81aeab79ccc46f75e645dd2252&oe=5A244C1F'),
  ('This is the second post', 'https://scontent.fmci1-1.fna.fbcdn.net/v/t31.0-8/19221825_10154823868368458_3888337127533159098_o.jpg?oh=3f7769b2f56b5d729ebc7fd9b9ff2c9f&oe=59F3AAF1');
