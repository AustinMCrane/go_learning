-- Schema for api
CREATE TABLE guest (
  id SERIAL PRIMARY KEY,
  last_name char varying(255),
  first_name char varying(255),
  wedding_party boolean default false,
  rsvp boolean default false
);

CREATE TABLE location (
  id SERIAL PRIMARY KEY,
  name char varying(255)
);

CREATE TABLE location_cordinate (
  id SERIAL PRIMARY KEY,
  location_id integer not null references location(id),
  lat float,
  lng float
);

CREATE TABLE location_marker (
  id SERIAL PRIMARY KEY,
  location_id integer not null references location(id),
  lat float,
  lng float,
  title char varying(64),
  image char varying(400),
  destination boolean default false
);

CREATE TABLE post (
  id SERIAL PRIMARY KEY,
  body char varying(255),
  image char varying(400)
);

CREATE TABLE timeline (
  id SERIAL PRIMARY KEY,
  body char varying(500),
  time char varying(100),
  image char varying(255)
);
