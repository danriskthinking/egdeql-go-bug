CREATE MIGRATION m1jc6xlr4trkgiejhdpgctsuxem6asovaaflea5q5w7anxvn6b63qa
    ONTO initial
{
  CREATE MODULE Test IF NOT EXISTS;
  CREATE TYPE Test::Test2;
  CREATE TYPE Test::Test {
      CREATE SINGLE LINK test_2: Test::Test2;
  };
};
