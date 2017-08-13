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

INSERT INTO timeline (body, time, image)
VALUES
  ('We met', 'December 29, 2012', 'https://scontent.fmci1-1.fna.fbcdn.net/v/t1.0-0/p206x206/62986_4720801491409_226450906_n.jpg?oh=f1819409882eb2a303be40a76c7525e8&oe=5A236C47'),
  ('We are in love', 'Febuary 2, 2012', 'https://scontent.fmci1-1.fna.fbcdn.net/v/t1.0-9/485215_4859203191365_1005380069_n.jpg?oh=83caa7909f2c1ca14885a43ebf7644e4&oe=5A2F3923'),
  ('We meet the family', 'March 3, 2012', 'https://scontent.fmci1-1.fna.fbcdn.net/v/t1.0-0/p206x206/969372_10151709309147182_1354214451_n.jpg?oh=61750b5a610e4b609b862a1b1ff980ca&oe=5A32F4CB'),
  ('Knocked her up', 'IDK', 'https://scontent.fmci1-1.fna.fbcdn.net/v/t1.0-9/10629756_10100511497966232_4531025860085571582_n.jpg?oh=87a336db702df14a041071d0ef8a8602&oe=5A3078E1'),
  ('Hadley is born', 'March 5, 2015', 'https://scontent.fmci1-1.fna.fbcdn.net/v/t1.0-9/10431543_10152923584118458_709911543536027567_n.jpg?oh=ebfb6e9bd9ab23b5e8fcf6522d7a6daf&oe=59F1D9F5'),
  ('We are engaged', 'IDK', 'https://scontent.fmci1-1.fna.fbcdn.net/v/t1.0-9/18057912_10154674134963458_280330944505678958_n.jpg?oh=4234c0bff65ef909e55184d1cb672806&oe=5A20EACF'),
  ('We are here today!', 'September 16, 2017', 'https://scontent.fmci1-1.fna.fbcdn.net/v/t1.0-9/19059264_10154823871378458_1487953269318551064_n.jpg?oh=07d66b274001220249369779a24a5172&oe=5A243662');

INSERT INTO location_cordinate (location_id, lat, lng)
VALUES
  (1, 3.34, 3.34),
  (1, 3.4, 3.4),
  (1, 3.5, 4),
  (1, 2, 6),
  (2, 10, 2),
  (2, 10.2, 3),
  (2, 20, 10);

INSERT INTO location_marker (location_id, lat, lng, title, image, destination)
VALUES
  (1, 3.34, 3.34, 'one', 'something', false),
  (1, 3.4, 3.4, 'two', 'something', false),
  (1, 3.5, 4, 'three', 'something', false),
  (1, 2, 6, 'four', 'something', true),
  (2, 4.44, 4.44, 'one', 'something', false),
  (2, 4.4, 4.4, 'two', 'something', false),
  (2, 4.5, 4, 'three', 'something', false),
  (2, 6, 6, 'four', 'something', true)
