import { test, expect } from "@playwright/test";
import dotenv from "dotenv"
dotenv.config()

test("should response one skill when request POST /api/v1/skills", async ({
	request,
}) => {
	const reps = await request.post("http://" + process.env.API_URL + ":8000" +"/api/v1/skills", {
		data: {
			key: "python",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});
	expect(reps.ok).toBeTruthy();
	const response = await reps.json();
	expect(response).toEqual(
		expect.objectContaining({
			message: "request accepted",
			status: "success",
		})
	);
	const getReps = await request.get("/api/v1/skills/python");
	const getResp = await getReps.json();
	expect(getResp).toEqual(
		expect.objectContaining({
			status: "success",
			data: expect.objectContaining({
				key: "python",
				name: "Python",
				description:
					"Python is an interpreted, high-level, general-purpose programming language.",
				logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
				tags: ["programming language", "scripting"],
			}),
		})
	);

	// const key = response["data"]["key"];
	await request.delete("/api/v1/skills/python");
});

test("should response with one skill when request GET /api/v1/skills/python3", async ({
	request,
}) => {
	const reps = await request.post("/api/v1/skills", {
		data: {
			key: "python3",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});
	expect(reps.ok).toBeTruthy();
	const response = await reps.json();
	expect(response).toEqual(
		expect.objectContaining({
			message: "request accepted",
			status: "success",
		})
	);

	const getReps = await request.get("/api/v1/skills/python3");
	expect(getReps.ok()).toBeTruthy();
	const getResponse = await getReps.json();
	expect(getResponse).toEqual(
		expect.objectContaining({
			status: "success",
			data: {
				key: "python3",
				name: "Python",
				description:
					"Python is an interpreted, high-level, general-purpose programming language.",
				logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
				tags: ["programming language", "scripting"],
			},
		})
	);

	const key = getResponse["data"]["key"];
	await request.delete("/api/v1/skills/" + key);
});

test("should response with all skill when request GET /api/v1/skills", async ({
	request,
}) => {
	const reps1 = await request.post("/api/v1/skills", {
		data: {
			key: "pythontest1",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});
	const reps2 = await request.post("/api/v1/skills", {
		data: {
			key: "pythontest2",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});
	const reps3 = await request.post("/api/v1/skills", {
		data: {
			key: "pythontest3",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});

	expect(reps1.ok()).toBeTruthy();
	expect(reps2.ok()).toBeTruthy();
	expect(reps3.ok()).toBeTruthy();

	const getReps = await request.get("/api/v1/skills");
	expect(getReps.ok()).toBeTruthy();
	const getResponse = await getReps.json();
	expect(getResponse).toEqual(
		expect.objectContaining({
			status: "success",
			data: expect.arrayContaining([
				{
					key: "pythontest1",
					name: "Python",
					description:
						"Python is an interpreted, high-level, general-purpose programming language.",
					logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
					tags: ["programming language", "scripting"],
				},
				{
					key: "pythontest2",
					name: "Python",
					description:
						"Python is an interpreted, high-level, general-purpose programming language.",
					logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
					tags: ["programming language", "scripting"],
				},
				{
					key: "pythontest3",
					name: "Python",
					description:
						"Python is an interpreted, high-level, general-purpose programming language.",
					logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
					tags: ["programming language", "scripting"],
				},
			]),
		})
	);

	await request.delete("/api/v1/skills/" + "pythontest1");
	await request.delete("/api/v1/skills/" + "pythontest2");
	await request.delete("/api/v1/skills/" + "pythontest3");
});

test("should response with updated skill when PUT /api/v1/skills/{key} when key is available", async ({
	request,
}) => {
	const reps = await request.post("/api/v1/skills", {
		data: {
			key: "python10",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});
	expect(reps.ok).toBeTruthy();

	const updatedReps = await request.put("/api/v1/skills/python10", {
		data: {
			name: "Python 3",
			description:
				"Python 3 is the latest version of Python programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["data"],
		},
	});
	expect(updatedReps.ok()).toBeTruthy();
	const updateResponse = await updatedReps.json();
	expect(updateResponse).toEqual(
		expect.objectContaining({
			message: "request accepted",
			status: "success",
		})
	);

	const getReps = await request.get("/api/v1/skills/python10");
	expect(getReps.ok()).toBeTruthy();
	const getResponse = await getReps.json();
	expect(getResponse).toEqual(
		expect.objectContaining({
			status: "success",
			data: {
				key: "python10",
				name: "Python 3",
				description:
					"Python 3 is the latest version of Python programming language.",
				logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
				tags: ["data"],
			},
		})
	);

	await request.delete("/api/v1/skills/python10");
});

test("should response with updated skill name when PATCH /api/v1/skills/{key}/actions/nam when key is available", async ({
	request,
}) => {
	const reps = await request.post("/api/v1/skills", {
		data: {
			key: "python11",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});
	expect(reps.ok).toBeTruthy();

	const updatedReps = await request.patch(
		"/api/v1/skills/python11/actions/name",
		{
			data: {
				name: "Python 3",
			},
		}
	);
	expect(updatedReps.ok()).toBeTruthy();
	const updateResponse = await updatedReps.json();
	expect(updateResponse).toEqual(
		expect.objectContaining({
			message: "request accepted",
			status: "success",
		})
	);

	const getReps = await request.get("/api/v1/skills/python11");
	expect(getReps.ok()).toBeTruthy();
	const getResponse = await getReps.json();
	expect(getResponse).toEqual(
		expect.objectContaining({
			status: "success",
			data: {
				key: "python11",
				name: "Python 3",
				description:
					"Python is an interpreted, high-level, general-purpose programming language.",
				logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
				tags: ["programming language", "scripting"],
			},
		})
	);

	await request.delete("/api/v1/skills/python11");
});

test("should response with updated skill when PATCH /api/v1/skills/{key}/actions/description when key is available", async ({
	request,
}) => {
	const reps = await request.post("/api/v1/skills", {
		data: {
			key: "python13",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});
	expect(reps.ok).toBeTruthy();

	const updatedReps = await request.patch(
		"/api/v1/skills/python13/actions/description",
		{
			data: {
				description:
					"Python 3 is the latest version of Python programming language.",
			},
		}
	);
	expect(updatedReps.ok()).toBeTruthy();
	const updateResponse = await updatedReps.json();
	expect(updateResponse).toEqual(
		expect.objectContaining({
			message: "request accepted",
			status: "success",
		})
	);

	const getReps = await request.get("/api/v1/skills/python13");
	expect(getReps.ok()).toBeTruthy();
	const getResponse = await getReps.json();
	expect(getResponse).toEqual(
		expect.objectContaining({
			status: "success",
			data: {
				key: "python13",
				name: "Python",
				description:
					"Python 3 is the latest version of Python programming language.",
				logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
				tags: ["programming language", "scripting"],
			},
		})
	);
	await request.delete("/api/v1/skills/python13");
});

test("should response with updated skill's logo when PATCH /api/v1/skills/python14/actions/logo", async ({
	request,
}) => {
	const reps = await request.post("/api/v1/skills", {
		data: {
			key: "python14",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});
	expect(reps.ok).toBeTruthy();

	const updatedReps = await request.patch(
		"/api/v1/skills/python14/actions/logo",
		{
			data: {
				logo: "updated logo",
			},
		}
	);
	expect(updatedReps.ok()).toBeTruthy();
	const updateResponse = await updatedReps.json();
	expect(updateResponse).toEqual(
		expect.objectContaining({
			message: "request accepted",
			status: "success",
		})
	);

	const getReps = await request.get("/api/v1/skills/python14");
	expect(getReps.ok()).toBeTruthy();
	const getResponse = await getReps.json();
	expect(getResponse).toEqual(
		expect.objectContaining({
			status: "success",
			data: {
				key: "python14",
				name: "Python",
				description:
					"Python is an interpreted, high-level, general-purpose programming language.",
				logo: "updated logo",
				tags: ["programming language", "scripting"],
			},
		})
	);
	await request.delete("/api/v1/skills/python14");
});

test("should response with updated skill's tags when PATCH /api/v1/skills/python15/actions/tags", async ({
	request,
}) => {
	const reps = await request.post("/api/v1/skills", {
		data: {
			key: "python15",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});
	expect(reps.ok).toBeTruthy();

	const updatedReps = await request.patch(
		"/api/v1/skills/python15/actions/tags",
		{
			data: {
				tags: ["programming language", "data"],
			},
		}
	);
	expect(updatedReps.ok()).toBeTruthy();
	const updateResponse = await updatedReps.json();
	expect(updateResponse).toEqual(
		expect.objectContaining({
			message: "request accepted",
			status: "success",
		})
	);

	const getReps = await request.get("/api/v1/skills/python15");
	expect(getReps.ok()).toBeTruthy();
	const getResponse = await getReps.json();
	expect(getResponse).toEqual(
		expect.objectContaining({
			status: "success",
			data: {
				key: "python15",
				name: "Python",
				description:
					"Python is an interpreted, high-level, general-purpose programming language.",
				logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
				tags: ["programming language", "data"],
			},
		})
	);
	await request.delete("/api/v1/skills/python15");
});

test("should response success when request DELETE /api/v1/skills/{key} when key is valid", async ({
	request,
}) => {
	const reps = await request.post("/api/v1/skills", {
		data: {
			key: "python16",
			name: "Python",
			description:
				"Python is an interpreted, high-level, general-purpose programming language.",
			logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
			tags: ["programming language", "scripting"],
		},
	});
	expect(reps.ok).toBeTruthy();

	const deleteReps = await request.delete("/api/v1/skills/python16");
  expect(deleteReps.ok()).toBeTruthy()
  const deleteResponse = await deleteReps.json()
  expect(deleteResponse).toEqual(
    expect.objectContaining({
      status:"success",
      message:"request accepted"
    })
  )
	const getReps = await request.get("/api/v1/skills/python16");
	const getResp = await getReps.json();
	expect(getResp).toEqual(
		expect.objectContaining({
			status: "error",
			message: "Skill not found"
		})
	);

});
