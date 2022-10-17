CREATE ROLE "{{name}}" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}';
GRANT ALL ON ALL TABLES IN SCHEMA public TO "{{name}}";