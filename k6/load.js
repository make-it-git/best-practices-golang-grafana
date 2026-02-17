import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  scenarios: {
    steady: {
      executor: 'constant-arrival-rate',
      rate: 50,
      timeUnit: '1s',
      duration: '1d',
      preAllocatedVUs: 20,
    }
  },
  thresholds: {
    http_req_duration: ['p(95)<4000'],
    http_req_failed: ['rate<0.5']
  }
};

export default function () {
  query('http://localhost:8080');
  if (Math.random() < 0.8) {
    query('http://localhost:8081');
  }
}

function query(base) {
  http.get(`${base}/success`);
  http.get(`${base}/error`);
  http.get(`${base}/slow`);
}