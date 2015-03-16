create table pre_auth (
    pre_auth_id number,
    account_id number,
    auth_amount number(18,2),
    auth_notes varchar(200)
);

create table pre_auth_results (
    pre_auth_id number,
    pre_auth_results_id number,
    success varchar(1),
    response varchar(200)
);


